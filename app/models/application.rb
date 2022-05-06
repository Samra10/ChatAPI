class Application < ApplicationRecord

	has_secure_token :application_token

	validates :name, presence: true
end
