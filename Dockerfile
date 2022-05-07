FROM ruby:3.1.2

# Default directory

WORKDIR /chat-api


COPY Gemfile Gemfile.lock ./
#RUN gem install bundler
RUN bundle install
COPY . .


# Start the main process.
CMD []