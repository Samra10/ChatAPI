Rails.application.routes.draw do

  resources :applications, only: [:index, :create] do
  end
end
