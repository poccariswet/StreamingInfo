package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	nsq "github.com/bitly/go-nsq"
	mgo "gopkg.in/mgo.v2"
)

// MongoDB接続
var db *mgo.Session

func dialdb() error {
	var err error
	log.Println("MongoDBにダイアル中: localhost")
	db, err = mgo.Dial("localhost")
	return err
}

func closedb() {
	db.Close()
	log.Println("DB接続が閉じられました")
}

// db読み込み
type poll struct {
	Options []string
}

func loadOptions() ([]string, error) {
	var options []string
	iter := db.DB("ballots").C("idle").Find(nil).Iter()
	var p poll
	for iter.Next(&p) {
		options = append(options, p.Options...)
	}
	iter.Close()
	return options, iter.Err()
}

// NSQへのパブリッシュ
func publishVotes(votes <-chan string) <-chan struct{} {
	stopchan := make(chan struct{}, 1)
	pub, _ := nsq.NewProducer("localhost:4150", nsq.NewConfig())
	go func() {
		for vote := range votes {
			pub.Publish("votes", []byte(vote)) //投票内容をパブリッシュ
		}
		log.Println("Publisher: 停止中...")
		pub.Stop()
		log.Println("Publisher: 停止完了")
		stopchan <- struct{}{}
	}()
	return stopchan
}

func main() {
	var stoplock sync.Mutex
	stop := false
	stopChan := make(chan struct{}, 1)
	signalChan := make(chan os.Signal, 1)
	go func() {
		<-signalChan
		stoplock.Lock()
		stop = true
		stoplock.Unlock()
		log.Println("停止...")
		stopChan <- struct{}{}
		closeConn()
	}()
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	//db接続を開始
	if err := dialdb(); err != nil {
		log.Fatalln("MongoDBへのダイヤルに失敗:", err)
	}
	defer closedb()

	//処理を開始
	votes := make(chan string) //投票結果のためのチャネル
	publisherStoppedChan := publishVotes(votes)
	twitterStoppedChan := startTwitterStream(stopChan, votes)
	go func() {
		for {
			time.Sleep(1 * time.Minute)
			closeConn()
			stoplock.Lock()
			if stop {
				stoplock.Unlock()
				break
			}
			stoplock.Unlock()
		}
	}()
	<-twitterStoppedChan
	close(votes)
	<-publisherStoppedChan
}
