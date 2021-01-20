protoc --go_out=plugins=grpc:../services/user/packages/api/proto/users --go_opt=paths=source_relative users.proto \
--go_out=plugins=grpc:../services/messaging/proto/users --go_opt=paths=source_relative users.proto \

