# Имя приложения. Используется как имя итогового бинарного файла.
APP_NAME := shortener

# Путь к Go-пакету, в котором находится main.go.
# Именно этот пакет будет запускаться через go run и собираться через go build.
MAIN := ./cmd/shortener

# Директория, куда будет положен собранный бинарный файл.
BIN_DIR := ./cmd/shortener

# Полный путь к итоговому исполняемому файлу.
# В результате make build получится: ./cmd/shortener/shortener
BIN := $(BIN_DIR)/$(APP_NAME)

# Go-команда. Обычно просто go, но можно переопределить при необходимости.
GO := go

# Дополнительные флаги для go build.
# Например: make build GOFLAGS="-race"
GOFLAGS :=

# Номер итерации для автотестов.
# По умолчанию запускается TestIteration1.
# Пример: make autotest ITER=2
ITER ?= 1

# Список phony-целей.
# Они не создают файлы с такими именами, а являются командами Makefile.
.PHONY: run build autotest clean

# Запускает приложение из исходников без создания бинарного файла.
run:
	$(GO) run $(MAIN)

# Собирает исполняемый файл по пути:
# ./cmd/shortener/shortener
build:
	$(GO) build $(GOFLAGS) -o $(BIN) $(MAIN)

# Удаляет собранный бинарный файл.
clean:
	rm -f $(BIN)

# Собирает приложение и запускает shortenertest.
#
# make autotest        -> TestIteration1
# make autotest ITER=2 -> TestIteration2
# make autotest ITER=3 -> TestIteration3

# В Makefile символ $ экранируется как $$,
# поэтому $$ в конце нужен для регулярки:
# ^TestIteration1$
autotest: build
	shortenertest -test.v -test.run=^TestIteration$(ITER)$$ \
		-binary-path=$(BIN)
