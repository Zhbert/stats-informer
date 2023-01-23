VERSION="0.2.1"

help:
	@echo "Hello from stats-informer v.$(VERSION)!"

build:
	make clean
	test -e ./target || mkdir ./target
	cp -r ./static ./target
	cp -r ./templates ./target
	go build -o ./target/stats-informer-v.$(VERSION) app/main.go

clean:
	rm -rf ./target
	rm -rf ./bundles

run:
	test -e ./target || make build
	./target/stats-informer-v.$(VERSION)

build_bundle:
	make clean
	make build
	touch ./target/repos.conf
	mkdir ./bundles
	mv ./target ./stats-informer-v.$(VERSION)
	zip -r ./bundles/stats-informer-v.$(VERSION).zip ./stats-informer-v.$(VERSION)/*
	rm -rf ./stats-informer-v.$(VERSION)

docker_build:
	build -t zhbert/stats-informer:$(VERSION) .
