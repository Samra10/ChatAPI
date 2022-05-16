# applications-chats-api
API for creating chats in RoR & Golang

## Install
```bash
sudo docker-compose down && sudo docker-compose build && sudo docker-compose up
```
Make sure that port `3000` & '8080' is availble to allow the API to run on it.

### Rails API availble end points

```
- Applications end points

GET   /applications/
POST  /applications?name={name}
GET   /applications/{application_token}
PUT   /applications/{application_token}?name={name}

- Chats end points
GET   /applications/{application_token}/chats
GET   /applications/{application_token}/chats/{chat_number}

- Messages end points
GET   /applications/{application_token}/chats/{chat_number}/messages
GET   /applications/{application_token}/chats/{chat_number}/messages/{message_number}
GET   /applications/{application_token}/chats/{chat_number}/messages/search?keyword={keyword}
PUT   /applications/{application_token}/chats/{chat_number}/messages/{message_number}?body={message_body}
```

### Go API availble end points
```
- Applications end points

POST  /applications/{application_token}/chats/
POST  /applications/{application_token}/chats/{chat_number}/messages?body={message_body}
```
#### Examples

Let's start by create new application:
```
-Request
GET 'http://localhost:3000/applications?name=newApp'

# Response
{
  "name": "newApp",
  "created_at": "2022-01-01T01:10:00",
  "updated_at": "2022-01-01T01:10:00",
  "chat_count": 0
  "application_token": "bs2oCw18zFGVUJmCAmTFhuBN",
}
```
Then create new Chat:

 ```
POST 'http://localhost:8080/applications/bs2oCw18zFGVUJmCAmTFhuBN/chats'

# Response
{
  "number": 1,
  "application_token": "bs2oCw18zFGVUJmCAmTFhuBN"
}
 ```

Then create new message:
 ```
'{"body": "New Messages"}'  POST 'http://localhost:8080/applications/bs2oCw18zFGVUJmCAmTFhuBN/chats/1/messages'

# Response
{
  "number": 1,
  "chat_number":1,
  "application_token": "bs2oCw18zFGVUJmCAmTFhuBN"
}
 ```
