
VERSION=v0.1.76

gen:
	cd proto && \
	buf generate

build: 
	
	go mod tidy
	echo ${VERSION} > VERSION
	git add .
	git commit -m "update"
	git push origin main
	git tag ${VERSION}
	git push origin ${VERSION}
	go list -m github.com/pnocera/res-gomodel@${VERSION}
	
