package configs

const AppAPIUrl = "http://chat-api:3000/"
const ChatsRoute = "/applications/{application_token}/chats"
const MessagesRoute = "/applications/{application_token}/chats/{chat_number}/messages"

const RedisUrl = "redis:6379"
const ChatQueue = "chat"
const ChatWorker = "ChatWorker"
const MessagesQueue = "message"
const MessagesWorker = "MessageWorker"
