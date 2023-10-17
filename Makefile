# used for repetitive grpc commands
clean:
	rm grpc-server/proto/*video*.go 

# these files should be added at the root of project
# https://stackoverflow.com/questions/66168350/import-google-api-annotations-proto-was-not-found-or-had-errors-how-do-i-add
download:
	mkdir -p google/api
	curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto > google/api/annotations.proto
	curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto > google/api/http.proto

create-test:
	protoc --go_out=grpc-server --go-grpc_out=grpc-server grpc-server/proto/test.proto

	protoc -I . --grpc-gateway_out . \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    grpc-server/proto/test.proto

create-video:
	protoc --go_out=grpc-server --go-grpc_out=grpc-server grpc-server/proto/video.proto

	protoc -I . --grpc-gateway_out . \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    grpc-server/proto/video.proto

tidy:
	cd grpc-server && go mod tidy