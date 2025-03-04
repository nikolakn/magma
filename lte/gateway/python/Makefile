include $(MAGMA_ROOT)/lte/gateway/python/defs.mk
# Contains common targets for gateway python development
include $(MAGMA_ROOT)/orc8r/gateway/python/python.mk

# Set up some pattern rules
define setup_rules
	$(eval PSRC := $1)
	CLEAN_LIST += $(PSRC)_clean
	BUILD_LIST += $(PSRC)_build
	TEST_LIST += $(PSRC)_test
endef
_ := $(foreach python_src, $(PYTHON_SRCS), $(eval $(call setup_rules,$(python_src))))

buildenv: setupenv protos swagger $(BUILD_LIST) py_patches
$(BUILD_LIST): %_build:
	make -C $* install_egg

test_all: buildenv $(BIN)/nosetests $(BIN)/coverage $(TEST_LIST)
$(TEST_LIST): %_test:
	make -C $* .test

define run_unit_tests
    . $(PYTHON_BUILD)/bin/activate; cd $(2); $(PYTHON_BUILD)/bin/nosetests -s $(1) || exit 1
endef
define run_sudo_unit_tests
    . $(PYTHON_BUILD)/bin/activate; cd $(2); sudo $(PYTHON_BUILD)/bin/nosetests -s $(1) || exit 1
endef

# unit_tests (UT_PATH=unit_test_path|MAGMA_SERVICE=service_name)[DONT_BUILD_ENV=1]
unit_tests: $(TESTS)
ifndef MAGMA_SERVICE
ifndef UT_PATH
	@echo "usage: make unit_tests (UT_PATH=unit_test_path|MAGMA_SERVICE=service_name)[DONT_BUILD_ENV=1]"
	@exit 1
endif
endif
ifdef MAGMA_SERVICE
	$(eval TEST_PATH_LTE ?= $(shell grep -oP '[^\s]+(?=$(MAGMA_SERVICE))[^\s]+' $(MAGMA_ROOT)/lte/gateway/python/defs.mk))
	$(eval TEST_PATH_ORC8R ?= $(shell grep -oP '[^\s]+(?=$(MAGMA_SERVICE))[^\s]+' $(MAGMA_ROOT)/orc8r/gateway/python/defs.mk))
else ifdef UT_PATH
	$(eval TEST_PATH_LTE ?= $(UT_PATH))
endif
	@if [ ! -d "$(MAGMA_ROOT)/lte/gateway/python/$(TEST_PATH)" ]; then if [ ! -d "$(MAGMA_ROOT)/orc8r/gateway/python/$(TEST_PATH_ORC8R)" ]; then echo "no tests found" && exit 1; fi; fi
ifndef DONT_BUILD_ENV
	@$(MAKE) buildenv $(BIN)/nosetests $(BIN)/coverage
endif
ifdef MAGMA_SERVICE
	sudo service magma@$(MAGMA_SERVICE) stop
endif
	$(eval NON_SUDO_TESTS ?= $(patsubst %, \%%,$(TESTS)))
	$(eval M_SUDO_TESTS ?= $(patsubst %, \%%,$(SUDO_TESTS)))
	$(eval SELECTED_TESTS ?= $(patsubst %/, %\\/, $(TEST_PATH_LTE)))

	$(if $(strip $(NON_SUDO_TESTS)),$(eval SEL_TEST_PATH_LTE ?= $(filter $(NON_SUDO_TESTS), $(SELECTED_TESTS))))
	$(if $(strip $(M_SUDO_TESTS)),$(eval SEL_SUDO_TEST_PATH_LTE ?= $(filter $(M_SUDO_TESTS), $(SELECTED_TESTS))))

	$(if $(strip $(SEL_TEST_PATH_LTE)),$(call run_unit_tests,$(SEL_TEST_PATH_LTE), $(MAGMA_ROOT)/lte/gateway/python/))
	$(if $(strip $(SEL_SUDO_TEST_PATH_LTE)),$(call run_sudo_unit_tests,$(SEL_SUDO_TEST_PATH_LTE), $(MAGMA_ROOT)/lte/gateway/python/))
	$(if $(strip $(TEST_PATH_ORC8R)), $(call run_unit_tests, $(TEST_PATH_ORC8R), $(MAGMA_ROOT)/orc8r/gateway/python/))

coverage: $(BIN)/coverage
	$(BIN)/coverage report

$(BIN)/nosetests: install_virtualenv
	$(VIRT_ENV_PIP_INSTALL) -I nose==1.3.7

$(BIN)/coverage: install_virtualenv
	$(VIRT_ENV_PIP_INSTALL) "coverage>=6.1.2"

$(BIN)/pylint: install_virtualenv
	$(VIRT_ENV_PIP_INSTALL) pylint==2.12.2

$(BIN)/pep8: install_virtualenv
	# pylint doesn't cover all the pep8 style guidelines. Specifically,
	# E203, E301, E303, W203, W291, W292
	$(VIRT_ENV_PIP_INSTALL) pep8==1.7.0

# Disable W0511: todo warnings
# Disable R0903: Too few public methods
CHECK_CMD_PYLINT := find . -name '*.py' -exec $(BIN)/pylint --disable=R0903,W0511 {} +;
CHECK_CMD_PEP8 := find . -name '*.py' -exec $(BIN)/pep8 {} +;
check: buildenv $(BIN)/pylint $(BIN)/pep8
	$(CHECK_CMD_PEP8)
	$(CHECK_CMD_PYLINT)

clean: $(CLEAN_LIST)
	sudo rm -rf $(PYTHON_BUILD)/ .coverage
	sudo find . -name '*.pyc' -o -name '__pycache__' -prune -exec rm -rf {} \;
$(CLEAN_LIST): %_clean:
	make -C $* remove_egg
