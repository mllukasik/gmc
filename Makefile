RELEASE_EMAIL=tvpozytyw@gmail.com
RELEASE_NAME=mllukasik

run:
	@go run . $(ARGS)

compile:
	@go build

test:
	@go test ./...

version:
	@go run . --version

prepare-release:
	@git config --global user.email "$(RELEASE_EMAIL)"
	@git config --global user.name "$(RELEASE_NAME)"
	@./scripts/properties version build_date
	@git add .
	@git commit -m "prepare release"
	@version=$$(go run . --version); \
		echo $$version; \
		git tag $$version
	@git push
	@git push origin --tags

after-release:
	@git config --global user.email "$(RELEASE_EMAIL)"
	@git config --global user.name "$(RELEASE_NAME)"
	@./scripts/properties version build_date
	@git add .
	@git commit -m "prepare for next development iteration"
	@git push

