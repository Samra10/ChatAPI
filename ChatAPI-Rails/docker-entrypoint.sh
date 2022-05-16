#!/bin/bash

rm -f /ChatAPI-Rails/tmp/pids/server.pid

/usr/bin/wait-for-it.sh db:3306 -t 0
/usr/bin/wait-for-it.sh redis:6379 -t 0
/usr/bin/wait-for-it.sh elasticsearch:9200 -t 0

cd ChatAPI-Rails
bundle exec sidekiq &
rails db:create db:migrate db:seed

#exec bundle exec"$@"
bundle exec rails server -b 0.0.0.0 -p 3000