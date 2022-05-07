class Message < ApplicationRecord

	belongs_to :chat, counter_cache: :messages_count, touch: true


	validates :number, presence: true, uniquness: { scope: :chat_id}

	include Searchable 
end
