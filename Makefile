PACKAGE_NAME := MineServerTools
SHELL := /bin/bash

.PHONY: package-deb
package-deb:
	if [ -f package/DEBIAN/control ]; then rm package/DEBIAN/control; fi
	if [ -f package/usr/bin/bedrock-tools/tools/download-server ]; then rm package/usr/bin/bedrock-tools/tools/download-server; fi
	mkdir -p dist
	mkdir -p package/DEBIAN
	mkdir -p package/usr/bin/bedrock-tools/tools
	cp control.template package/DEBIAN/control
	go build -o package/usr/bin/bedrock-tools/tools/update-bedrock tools/bedrock/update/*.go
	go build -o package/usr/bin/bedrock-tools/tools/backup-bedrock tools/bedrock/backup/*.go
	go build -o package/usr/bin/bedrock-tools/tools/info-bedrock tools/bedrock/system/system.go
	sed -i "s/x.y.z/$(CURRENT_VERSION_MICRO)/" package/DEBIAN/control
	dpkg-deb --build package/ dist/$(PACKAGE_NAME)_$(CURRENT_VERSION_MICRO)_all.deb
	if [ -f package/DEBIAN/control ]; then rm package/DEBIAN/control; fi
	if [ -f package/usr/bin/bedrock-tools/tools/backup-bedrock ]; then rm package/usr/bin/bedrock-tools/tools/backup-bedrock; fi
	if [ -f package/usr/bin/bedrock-tools/tools/info-bedrock ]; then rm package/usr/bin/bedrock-tools/tools/info-bedrock; fi
	if [ -f package/usr/bin/bedrock-tools/tools/update-bedrock ]; then rm package/usr/bin/bedrock-tools/tools/update-bedrock; fi
## Gerenciamento de versões

MAKE               := make --no-print-directory

DESCRIBE           := $(shell git describe --match "v*" --always --tags)
DESCRIBE_PARTS     := $(subst -, ,$(DESCRIBE))

VERSION_TAG        := $(word 1,$(DESCRIBE_PARTS))
COMMITS_SINCE_TAG  := $(word 2,$(DESCRIBE_PARTS))

VERSION            := $(subst v,,$(VERSION_TAG))
VERSION_PARTS      := $(subst ., ,$(VERSION))

MAJOR              := $(word 1,$(VERSION_PARTS))
MINOR              := $(word 2,$(VERSION_PARTS))
MICRO              := $(word 3,$(VERSION_PARTS))

NEXT_MAJOR         := $(shell echo $$(($(MAJOR)+1)))
NEXT_MINOR         := $(shell echo $$(($(MINOR)+1)))
NEXT_MICRO          = $(shell echo $$(($(MICRO)+$(COMMITS_SINCE_TAG))))

ifeq ($(strip $(COMMITS_SINCE_TAG)),)
CURRENT_VERSION_MICRO := $(MAJOR).$(MINOR).$(MICRO)
CURRENT_VERSION_MINOR := $(CURRENT_VERSION_MICRO)
CURRENT_VERSION_MAJOR := $(CURRENT_VERSION_MICRO)
else
CURRENT_VERSION_MICRO := $(MAJOR).$(MINOR).$(NEXT_MICRO)
CURRENT_VERSION_MINOR := $(MAJOR).$(NEXT_MINOR).0
CURRENT_VERSION_MAJOR := $(NEXT_MAJOR).0.0
endif

DATE                = $(shell date +'%d.%m.%Y')
TIME                = $(shell date +'%H:%M:%S')
COMMIT             := $(shell git rev-parse HEAD)
AUTHOR             := $(firstword $(subst @, ,$(shell git show --format="%aE" $(COMMIT))))
BRANCH_NAME        := $(shell git rev-parse --abbrev-ref HEAD)

TAG_MESSAGE         = "$(TIME) $(DATE) $(AUTHOR) $(BRANCH_NAME)"
COMMIT_MESSAGE     := $(shell git log --format=%B -n 1 $(COMMIT))

CURRENT_TAG_MICRO  := "v$(CURRENT_VERSION_MICRO)"
CURRENT_TAG_MINOR  := "v$(CURRENT_VERSION_MINOR)"
CURRENT_TAG_MAJOR  := "v$(CURRENT_VERSION_MAJOR)"

# --- Version commands ---

.PHONY: version
version:
	@$(MAKE) version-micro

.PHONY: version-micro
version-micro:
	@echo "$(CURRENT_VERSION_MICRO)"

.PHONY: version-minor
version-minor:
	@echo "$(CURRENT_VERSION_MINOR)"

.PHONY: version-major
version-major:
	@echo "$(CURRENT_VERSION_MAJOR)"

# --- Tag commands ---

.PHONY: tag-micro
tag-micro:
	@echo "$(CURRENT_TAG_MICRO)"

.PHONY: tag-minor
tag-minor:
	@echo "$(CURRENT_TAG_MINOR)"

.PHONY: tag-major
tag-major:
	@echo "$(CURRENT_TAG_MAJOR)"

# -- Meta info ---

.PHONY: tag-message
tag-message:
	@echo "$(TAG_MESSAGE)"

.PHONY: commit-message
commit-message:
	@echo "$(COMMIT_MESSAGE)"