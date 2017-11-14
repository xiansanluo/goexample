LIULISHUO_PROTOS_PATH = ./engzo-protos/protos
GOOGLE_API_PATH = ./engzo-protos/googleapis

vpath %.proto $(LIULISHUO_PROTOS_PATH)

make:  get

vpath %.proto $(LIULISHUO_PROTOS_PATH)

test:
	cd handler && go test

get:
	https_proxy=${HTTPS_PROXY} glide install

gen-golang:
	git submodule update --init
	protoc -I $(LIULISHUO_PROTOS_PATH) \
		-I $(GOOGLE_API_PATH) \
		--go_out=plugins=grpc:${GOPATH}/src \
		$(LIULISHUO_PROTOS_PATH)/liulishuo/backend/llspay/*.proto
	protoc -I $(LIULISHUO_PROTOS_PATH) \
		-I $(GOOGLE_API_PATH) \
		--go_out=plugins=grpc:${GOPATH}/src \
		$(LIULISHUO_PROTOS_PATH)/liulishuo/backend/llspay/grpc/*.proto
