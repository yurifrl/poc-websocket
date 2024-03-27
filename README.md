


protoc --go_out=. --go_opt=paths=source_relative proto/message.proto


docker compose run --rm app protoc --go_out=. --go_opt=paths=source_relative proto/message.proto