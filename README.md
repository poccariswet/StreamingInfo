mongo db & nsq command

```sh
$ nsqlookupd
```
```sh
$ nsqd --lookupd-tcp-address=localhost:4160
```
```sh
$ mongod --dbpath ./db
```
```sh
$ nsq_tail --topic="votes" --lookupd-http-address=localhost:4161
```
<br><br>

###mongo
Create a new poll in the ballots database:

For example

```sh
> use ballots
> db.polls.insert({title:"My poll",options:["one","two","three"]})
```

After a while, see the results by printing the polls:

```sh
> db.polls.find().pretty()
```

Delete the database of polls:

```sh
> db.polls.drop();
```