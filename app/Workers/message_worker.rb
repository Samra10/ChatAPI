class MessageWorker
  include Sidekiq::Worker
  sidekiq_options queue: :message

  def perform(application_token, chat_number, number, messageBody)
    application = Application.find_by!(application_token: application_token)
    chat = application.chats.find_by!(number: chat_number)
    chat.messages.create!(number: number, messageBody: message_body)
  end
end