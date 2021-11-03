debug-build:
	go build -o ./bin/terraform_provider_pfptmeta_linux_amd64 -gcflags="all=-N -l"
	./bin/terraform_provider_pfptmeta_linux_amd64 --debug

mod-tidy:
	go mod tidy


unittest:
	go test ./... $(if $(VERBOSE),"-v") -timeout 120m


acc_tests:
	TF_ACC=1 go test ./... $(if $(VERBOSE),"-v") -run "TestAcc*" -timeout 120m


generate:
	go generate -v -x


tests: acc_tests unittest


verify_clean: mod-tidy generate
	! git status -s | grep "??" || (echo "Uncommitted files found" && exit 1)
	git diff --stat --exit-code || (echo "Uncommitted files found" && exit 1)

local-release: verify_clean tests
	gpg --batch --import $(GPG_SECRET_PATH) && goreleaser release --rm-dist --snapshot


release: verify_clean tests
	gpg --batch --import $(GPG_SECRET_PATH) && goreleaser release --rm-dist