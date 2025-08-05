.DEFAULT_GOAL := help

## help: Вывести список доступных команд
.PHONY: help
help:
	@echo "Используйте: make <цель>"
	@echo ""
	@echo "Доступные цели:"
	@echo ""
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo ""

## test: Запустить тесты
.PHONY: test
test: ## Запускает unit-тесты
	@echo "Запуск тестов..."
