COVERAGE_DIR ?= .coverage.out

test:
	@if [ -d $(COVERAGE_DIR) ]; then rm -r $(COVERAGE_DIR); fi
	@mkdir $(COVERAGE_DIR)
	make test-with-flags TEST_FLAGS='-v -race -covermode atomic -coverprofile $$(COVERAGE_DIR)/combined.txt -bench=. -benchmem -timeout 20m'

test-with-flags:
	@go test $(TEST_FLAGS) ./...

# Proto-Service auto sync: generate proto code and sync service implementations
# Proto-Service 自动同步：生成 proto 代码并同步服务实现
orz:
	cd zhipin-kratos && make orz
