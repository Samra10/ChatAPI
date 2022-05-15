class MessagesController < ApplicationController

  include Render

  before_action :get_application
  before_action :get_chat

  def index
    @messages = @chat.messages.all
    render_json @messages
  end

  def show
    @message = @chat.messages.find_by!(number: params[:number])
    render_json @message
  end

  def search
    render_json Message.search(params[:keyword], @chat.id)
  end

  

  private

  def message_params
    params.permit(:messageBody)
  end

  def get_application
    @application = Application.find_by!(application_token: params[:application_token)
  end

  def set_chat
    @chat = @application.chats.find_by!(number: params[:chat_number])
  end
end
