class AddIndexToApplicationToken < ActiveRecord::Migration[7.0]
  def change
    add_index :applications, :application_token, unique: true
  end
end
