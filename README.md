# applications-chats-api
API for creating chats in RoR.

## Install
```bash
sudo docker-compose down && sudo docker-compose build && sudo docker-compose up
```
Make sure that port `3000` is availble to allow the API to run on it.

### API availble end points

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

#### Examples

Let's start by create new application:
```
-Request
'http://localhost:3000/applications?name=newApp'

Response
{
  "name": "newApp",
  "created_at": "2022-01-01T01:10:00",
  "updated_at": "2022-01-01T01:10:00",
  "chat_count": 0
  "application_token": "dsa585dahjhjqwkssasda",
}
```
Then create new Chat:
- [ ] ToDo/ Add Golang endpoint for creating new messages
- ```POST  /applications/{application_token}/chats/```

Then create new message:
- [ ] ToDo/ Add Golang endpoint for creating new messages
- ```POST  /applications/{application_token}/chats/{chat_number}/messages?body={messageBody}```


