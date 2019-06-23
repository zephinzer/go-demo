include ../../scripts/Makefile

define HELP_OUTPUT

Vault by Hashicorp
^^^^^^^^^^^^^^^^^^
Variant: with auto-unseal

Instructions:

0. Preperations
   - go to aws console and create a new iam user
   - ensure iam user has programmatic access
   - ensure iam user has AWSKeyManagementServicePowerUser
	 - note the access key and secret key of the iam user
   - go to kms and create a key, note the key id
1. Run `make init` to setup your AWS credentials
2. Run `make play` to start the setup
3. Run `make play_status` to checkout vault's status
4. Run `make play_init` to initialise vault
5. Run `make play_status` to checkout an unsealed vault
   - vault here was unsealed by using the kms
   - you should see both initialized and unsealed set to true
6. Run `make play_exec` to gain a shell/do whatever you want
7. Run `make play_stop` to tear down this shiz

Have fun!

endef
export HELP_OUTPUT

# sets your variables in the file at ./resources/config.hcl
init:
	#
	# set up kms details
	#
	# enter your aws region...
	read region && sed -e "s|__region__|\"$${region}\"|g" ./resources/config.template.hcl | tee ./resources/config.hcl
	# enter your iam access key...
	read access_key && sed -i -e "s|__access_key__|\"$${access_key}\"|g" ./resources/config.hcl
	# enter your iam secret key...
	read secret_key && sed -i -e "s|__secret_key__|\"$${secret_key}\"|g" ./resources/config.hcl
	# enter your kms key id...
	read kms_key_id && sed -i -e "s|__kms_key_id__|\"$${kms_key_id}\"|g" ./resources/config.hcl

# sets up the environment
play:
	@if ! [ -f "./resources/config.hcl" ]; then \
		printf -- '> use "make init" to generate the ./resources/config.hcl first\n\n'; \
		exit 1; \
	fi
	@docker-compose down
	@docker-compose up -d -V

# use this to check the status of vault
play_status:
	@$(MAKE) _vault CMD="vault status"

# use this to initialise vault
play_init:
	@$(MAKE) _vault CMD="vault operator init \
		-recovery-shares=3 \
		-recovery-threshold=2"
	# copy the recovery keys out
	# also copy the root token for use with your login
	# note that vault is initialised and unsealed (use `play_status` to see again)

# use this to login to vault via the cli (use root token from `make play_init`)
play_login:
	# enter the Root Token from `play_init` here
	@$(MAKE) _vault CMD="vault login"

# just give me a damn shell
play_exec:
	@$(MAKE) _vault ARGS="-it" CMD="/bin/sh"

# tears down the environment
play_stop:
	@docker-compose down

_vault:
	@docker exec ${ARGS} go_demo_init_vault_auto_unseal ${CMD}
