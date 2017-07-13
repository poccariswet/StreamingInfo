mongo db & nsq command

[1] start nsqlookupd demon
```sh
$ nsqlookupd
```

[2] start nsqd demon
```sh
$ nsqd --lookupd-tcp-address=localhost:4160
```

[3] start MongoDB demon
```sh
$ mongod --dbpath ./db
```

[4] build counter and starting in counter folder
```sh
$ go build -o counter
$ ./counter
```

[5] build twittervotes and starting in twittervotes folder
```sh
$ go build -o twittervotes
$ ./twittervotes
```

[6] build api and starting in api folder
```sh
$ go build -o api
$ ./api
```

[7] build web and starting in web folder
```sh
$ go build -o web
$ ./web
```

connect nsq and output all specific topic you set
```sh
$ nsq_tail --topic="votes" --lookupd-http-address=localhost:4161
```
<br><br>

#### mongoDB shell
Create a new poll in the ballots database:

For example

add the polls collection new topic
```sh
> use ballots
> db.polls.insert({"title":"My poll","options":["one","two","three"]})
```

After a while, see the results by printing the polls:

```sh
> db.polls.find().pretty()
```

Delete the database of polls:

```sh
> db.polls.drop();
```
