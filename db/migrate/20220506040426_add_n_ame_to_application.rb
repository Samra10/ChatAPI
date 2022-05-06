class AddNAmeToApplication < ActiveRecord::Migration[7.0]
  def change
    add_column :applications, :name, :string
  end
end
