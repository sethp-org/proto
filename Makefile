--clear-gen-proto:
	rm -rf internal
--gen-proto:
	mkdir -p internal
	docker run --rm --platform linux/amd64 -v $(shell pwd)/proto:/defs namely/protoc-all -d . -l go -o internal
--move-proto:
	mv proto/internal internal
	mv internal/internal/* internal
	mv internal/github.com/sethp-org/proto/internal/* internal
	rm -rf internal/internal
	rm -rf internal/github.com

gen: --clear-gen-proto --gen-proto --move-proto