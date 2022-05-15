module Response
	def response_json(object, status = :ok)
		render :json => object, :except => [:id, :application_id], status: status
		
	end
end