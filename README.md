# GO Lang REST API Excercise

```
    4. Please describe what metrics you would monitor to track the performance of the chat server.
    -> I would monitor the mem usage . Because in this application , all the operation I/O mainly happen on the mem cache . When the data grows bigger and the data operations becomes more frequent , the mem usage will be high .
    
    5. Bonus Points: please describe how you would improve the chat server API.
    -> ElasticCache cluster and Database cluster would help to reduce the huge I/O load from chat server .
```

## Quick Start


``` bash
# Install mux router
go get -u github.com/gorilla/mux
```

``` bash
go build
./go-restful-api
```

## Endpoints

### List 100 most recent messages
``` bash
curl -X GET -H "Content-Type: application/json" http://localhost:8081/messages

#Expect Output
#[{"timestamp":"1491345710.18","user":"superman","text":"hello"},
#{"timestamp":"1491345713.18","user":"batman","text":"hello"}]
```
### A request to post the given message
``` bash
curl -X POST -H "Content-Type: application/json" --data '{"user":"ironman", "text":"hello"}' http://localhost:8081/messages

#Expect Output
#{"ok":true}
```

### A request to return a set of users that have posted messages so far.
``` bash
curl -H "Content-Type: application/json" http://localhost:8081/users

#Expect Output
#["superman","batman"]
```
