class AddChatCountToApplication < ActiveRecord::Migration[7.0]
  def change
    add_column :applications, :chats_count, :integer
    add_column :applications, :application_token, :string
  end
end
