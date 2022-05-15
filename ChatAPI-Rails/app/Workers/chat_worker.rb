class ChatWorker
  include Sidekiq::Worker
  sidekiq_options queue: :chat

  def perform(appllication_token, number)
    application = Application.find_by!(application_token: appllication_token)
    application.chats.create!(number: number, messages_count: 0)
  end
end