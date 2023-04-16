--clear-gen-proto:
	rm -rf pb
--gen-proto:
	mkdir -p pb
	docker run --rm --platform linux/amd64 -v $(shell pwd)/proto:/defs namely/protoc-all -d . -l go -o pb
--move-proto:
	mv proto/pb pb
	mv pb/pb/* pb
	mv pb/github.com/sethp-org/proto/pb/* pb
	rm -rf pb/pb
	rm -rf pb/github.com

gen: --clear-gen-proto --gen-proto --move-proto