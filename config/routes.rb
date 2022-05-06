Rails.application.routes.draw do

  resources :applications, param: :application_token, only: [:index, :create, :show, :update] do
  end
end
