class ApplicationsController < ApplicationController

  include Response

  def index
    @applications = Application.all
    response_json @applications
  end

  def create
    @application = Application.new(application_params)
    @application.chats_count = 0

    @application.save
    response_json @application
  end

  def show
    @application = Application.find_by!(application_token: params[:application_token])
    response_json @application
  end

  def update
    @application = Application.find_by!(application_token: params[:application_token])
    @application.update(application_params)
    response_json @application
  end



  private

  def application_params
    params.permit(:name)
  end
end
