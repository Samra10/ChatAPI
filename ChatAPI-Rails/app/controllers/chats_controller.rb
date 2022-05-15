class ChatsController < ApplicationController

  include Render
  before_action :get_application

  def index
    @chats = @application.chats.all
    render_json @chats
  end

  def show
    @chat = @application.chats.find_by!(number: params[:number])
  end



  private

  def chat_params
    params.permit(:number)
  end

  def get_application
    @application = Application.find_by!(application_token: params[:application_token)
  end
end
