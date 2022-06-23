module example.com/hello

go 1.18

replace example.com/greetings/greetings => ../greetings

require example.com/greetings/greetings v1.1.0
