SI_VERSION=`cat $(PWD)/SI_VERSION`
SI_DIR := $(PWD)/assets/simple-icons
SI_RELEASE := $(SI_DIR)/$(SI_VERSION).zip

clean:
	rm -Rf $(PWD)/assets/simple-icons/*

source:
	curl -Ls https://github.com/simple-icons/simple-icons/archive/refs/tags/$(SI_VERSION).zip -z $(SI_RELEASE) -o $(SI_RELEASE)
	unzip $(SI_RELEASE) -d $(SI_DIR)
	rm $(SI_RELEASE)

test:
	go test ./... -v

cover:
	go test ./... -coverpkg=./... -coverprofile cp.out
	go tool cover -html=cp.out