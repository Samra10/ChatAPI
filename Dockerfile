FROM ruby:3.1.2-alpine

# Default directory
RUN mkdir /chat-api 
WORKDIR /chat-api
COPY Gemfile Gemfile.lock ./
RUN gem install bundler
RUN bundle check || bundle install
COPY . .

COPY docker-entrypoint.sh /usr/bin/
RUN chmod +x /usr/bin/docker-entrypoint.sh

ENTRYPOINT [ "docker-entrypoint.sh" ]
EXPOSE 3000

# Start the main process
CMD ["rails", "server", "-b", "0.0.0.0", "-p", "3000"]