# changeproto


To build proto files

DB
``` bash
protoc.exe -I db_service\ db_service\db.proto --go_out=plugins=grpc:db_service  
```                                         

Plaid
``` bash
protoc.exe -I plaid_service\ plaid_service\plaid.proto --go_out=plugins=grpc:plaid_service
```  
