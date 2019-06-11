SYS_OS :=
SYS_ARCH :=
BINEXT :=

ifeq ($(OS),Windows_NT)
	# this is windows
	SYS_OS = windows

	# get SYS_ARCH
	ifeq ($(PROCESSOR_ARCHITEW6432),AMD64)
		SYS_ARCH = amd64
	else
		ifeq ($(PROCESSOR_ARCHITECTURE),AMD64)
			SYS_ARCH = amd64
		endif
		ifeq ($(PROCESSOR_ARCHITECTURE),x86)
				SYS_ARCH = 386
		endif
	endif
else
	# get SYS_OS
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		SYS_OS = linux
	endif
	ifeq ($(UNAME_S),Darwin)
		SYS_OS = darwin
	endif
	# get SYS_ARCH
	UNAME_P := $(shell uname -p)
	ifeq ($(UNAME_P),x86_64)
		SYS_ARCH = amd64
	endif
	ifneq ($(filter %86,$(UNAME_P)),)
		SYS_ARCH = 386
	endif
	ifneq ($(filter arm%,$(UNAME_P)),)
		SYS_ARCH = arm
	endif
endif

sys:
	@printf -- "SYS_OS   : '$(SYS_OS)'\n"
	@printf -- "SYS_ARCH : '$(SYS_ARCH)'\n"
	@printf -- "BIN_EXT  : '$(BIN_EXT)'\n"
