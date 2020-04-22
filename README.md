# changeproto


To build proto files

``` bash
protoc.exe -I db_service\ db_service\db.proto --go_out=plugins=grpc:db_service  
```                                         
