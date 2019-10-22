require 'json'
require 'sinatra' 
configure { 
  set :server, :puma 
  set :port, 5000
}

post '/' do
  req = JSON.parse(request.body.read)
  
  req_text = req['request']['original_utterance']
  resp_text = req_text.length == 0 ? 'hello' : req_text
  
  content_type :json
  { :response => {
      :text => resp_text
    }, 
    :session => req['session'],
    :version =>  req['version'] }.to_json
end