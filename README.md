# go-microservices
Entire codebase is meant for golang practice and building microservices using Golang

## Youtube Playlist Link
https://youtube.com/playlist?list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_&si=WErbZ-ktbhVNfJMV

## Swagger Execution for reference:
1) go install github.com/go-swagger/go-swagger/cmd/swagger@latest => This gets all the binaries required for swagger
2) swagger generate spec -o ./swagger.yaml --scan-models => Generates the swagger.yaml file
3) I have a shortcut created generate-swagger.bat ! Execute: ./generate-swagger to recreate swagger.yaml file
4) Run your main function and Check your localhost:<"port no">/docs for the documentation 