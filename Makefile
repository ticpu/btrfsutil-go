all:
	c-for-go -ccincl btrfsutil.yaml

clean:
	rm -f btrfsutil/cgo_helpers.go btrfsutil/cgo_helpers.h btrfsutil/doc.go btrfsutil/types.go btrfsutil/const.go
	rm -f btrfsutil/btrfsutil.go

test:
	cd btrfsutil && go build
