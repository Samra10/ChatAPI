class Chat < ApplicationRecord

	belongs_to :application, counter_cache: :chat_count, touch: true

	has_many :messages

	validates :number, presence: true, uniquness: { scope: :application_id}
end
