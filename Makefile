gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb

build: 
	export VERSION=v0.1.68
	go mod tidy
	echo $(VERSION) > VERSION
	git add .
	git commit -m "update"
	git push origin main
	git tag $(VERSION)
	git push origin $(VERSION)
	go list -m github.com/pnocera/res-gomodel@$(VERSION)
	
