all:
	@protoc.exe -I db_service\ db_service\db.proto --go_out=plugins=grpc:db_service
	@protoc.exe -I plaid_service\ plaid_service\plaid.proto --go_out=plugins=grpc:plaid_service
	@protoc.exe -I stripe_service\ stripe_service\stripe.proto --go_out=plugins=grpc:stripe_service
	@protoc.exe -I mail_service\ mail_service\mail.proto --go_out=plugins=grpc:mail_service

db:
	@protoc.exe -I db_service\ db_service\db.proto --go_out=plugins=grpc:db_service

plaid:
	@protoc.exe -I plaid_service\ plaid_service\plaid.proto --go_out=plugins=grpc:plaid_service

stripe:
	@protoc.exe -I stripe_service\ stripe_service\stripe.proto --go_out=plugins=grpc:stripe_service

mail:
	@protoc.exe -I mail_service\ mail_service\mail.proto --go_out=plugins=grpc:mail_service
