module Searchable
  extend ActiveSupport::Concern

  included do
    include Elasticsearch::Model
    include Elasticsearch::Model::Callbacks

    settings do
      mapping dynamic: false do
        indexes :messageBody, type: :text, analyzer: :english
        indexes :chat_id
      end
    end

    def as_indexed_json(options={})
      self.as_json(only: [:messageBody, :number, :chat_id])
    end

    def self.search(term, chat_id)
      response = __elasticsearch__.search(
          query: {
              bool: {
                  must: [
                      { match: { chat_id: chat_id } },
                      { query_string: { query: "*#{term}*", fields: [:messageBody] } }
                  ]
              }
          }
      )
      response.results.map { |r| {messageBody: r._source.messageBody, number: r._source.number} }
    end
  end
end