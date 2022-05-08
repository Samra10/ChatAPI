#!/bin/bash

rm -f /chat-api/tmp/pids/server.pid

/usr/bin/wait-for-it.sh db:3306 -t 0
/usr/bin/wait-for-it.sh redis:6379 -t 0
/usr/bin/wait-for-it.sh elasticsearch:9200 -t 0
/usr/bin/wait-for-it.sh db:3306 -t 0

cd chat-api
bundle exec sidekiq &
rails db:create db:migrate db:seed

exec"$@"