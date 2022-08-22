# mage_test_case
Mage Games Backend Case Study

#### testing host
```sh
http://139.59.213.250:8080
```

## register
```sh
Post /v1/user/register
```
### params

```javascript
{
  "username" : string,required,alphanum,min=3,max=10
  "password" : bool,required,alphanum,min=3,max=10
}
```

### example response

```javascript
{
"status": "success",
"timestamp": 1661201326,
"result":{
  "id": 18,
  "username": "erdem12",
  "password": "5375Erdem"
  }
}
```

## login
```sh
Post /v1/user/login
```
### params

```javascript
{
  "username" : string,required,alphanum,min=3,max=10
  "password" : bool,required,alphanum,min=3,max=10
}
```

### example response

```javascript
{
"status": "success",
"timestamp": 1661201326,
"result":{
  "id": 18,
  "username": "erdem12"
  }
}
```


## endgame
```sh
Post /v1/endgame
```
### params

```javascript
{
"players":
   [
    {"id":13, "score":15}, 
    {"id":14, "score":20},  
    {"id":15, "score":9}
   ]
 }
```

### example response

```javascript
{
"status": "success",
"timestamp": 1661201444,
"result":[
    {
      "id": 16,
      "score": 25
    },
    {
      "id": 15,
      "score": 45
    },
    {
      "id": 13,
      "score": 75
    },
    {
      "id": 14,
      "score": 100
    }
  ]
}
```


## get leaderboard
```sh
Post /v1/leaderboard
```
### params

```javascript
-
```

### example response

```javascript
{
  "status": "success",
  "timestamp": 1661201483,
  "result": [
    {
      "id": 14,
      "rank": 12
    },
    {
      "id": 13,
      "rank": 9
    },
    {
      "id": 15,
      "rank": 6
    },
    {
      "id": 16,
      "rank": 3
    },
    {
      "id": 18,
      "rank": 0
    }
  ]
}
```
