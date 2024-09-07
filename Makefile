PACKAGE_NAME := MineServerTools
SHELL := /bin/bash

.PHONY: package-deb
package-deb:
	mkdir -p dist
	mkdir -p /tmp/deb-mineservertools/usr/bin/ /tmp/deb-mineservertools/usr/lib/systemd/system /tmp/deb-mineservertools/DEBIAN
	mkdir -p /tmp/deb-mineservertools/etc/mineservertools /tmp/deb-mineservertools/var/log /tmp/deb-mineservertools/var/mine-backups
	cp App/debug/debian/debian.control /tmp/deb-mineservertools/DEBIAN/control
	cp App/debug/debian/install.sh /tmp/deb-mineservertools/DEBIAN/postinst
	cp App/bin/bedrock/systemd/* /tmp/deb-mineservertools/usr/lib/systemd/system/
	cp App/bin/bedrock/shell/* /tmp/deb-mineservertools/usr/bin/
	cp App/bin/bedrock/logs/* /tmp/deb-mineservertools/var/log/
	cp App/bin/bedrock/confs/* /tmp/deb-mineservertools/etc/mineservertools/
	go build -o /tmp/deb-mineservertools/usr/bin/bed-tools App/bin/bedrock/*.go
	go build -o /tmp/deb-mineservertools/usr/bin/mtools-api App/bin/api/*.go
	sed -i "s/x.y.z/$(CURRENT_VERSION_MICRO)/" /tmp/deb-mineservertools/DEBIAN/control
	dpkg-deb --build /tmp/deb-mineservertools/ dist/$(PACKAGE_NAME)_$(CURRENT_VERSION_MICRO)_all.deb
	[ -d "/tmp/deb-mineservertools/" ] && rm -rf "/tmp/deb-mineservertools/"
	clear ; echo "Pacote DEBIAN criado com sucessso!"

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