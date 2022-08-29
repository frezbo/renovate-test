CONFORMANCE_IMAGE ?= ghcr.io/siderolabs/conform:latest

.PHONY: conformance
conformance:
	@docker pull $(CONFORMANCE_IMAGE)
	@docker run --rm -it -v $(PWD):/src -w /src $(CONFORMANCE_IMAGE) enforce
