

cd chat-api
bundle exec sidekiq &
rails db:create db:migrate db:seed

exec"$@"