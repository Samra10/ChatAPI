Rails.application.routes.draw do

  resources :applications, param: :application_token, only: [:index, :create, :show, :update] do
    resources :chats, param: :number, only: [:index, :show] do
      resources :messages, param: :number, only: [:index, :show] do
        collection do
          get :search
        end
      end
    end
  end
end
