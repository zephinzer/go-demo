include ../../scripts/Makefile

# sets up the environment
play:
	@docker-compose down
	@docker-compose up -d -V

# tears down the environment
play_stop:
	@docker-compose down

# use this to initialise vault
play_init:
	@$(MAKE) _vault CMD="vault operator init \
		-key-shares=3 \
		-key-threshold=2"
	# copy the Unseal Keys and Root Token and paste it somewhere!

# just give me a damn shell
play_exec:
	@$(MAKE) _vault CMD="/bin/sh"

# use this to check the status of vault
play_status:
	@$(MAKE) _vault CMD="vault status"

# use this to unseal vault (use keys from `make play_init`)
play_unseal:
	# run this 3 times using a different Unseal Key from `play_init` each time
	@$(MAKE) _vault CMD="vault operator unseal"

# use this to login to vault via the cli (use root token from `make play_init`)
play_login:
	# enter the Root Token from `play_init` here
	@$(MAKE) _vault CMD="vault login"

# use this to enable the aws dynamic secrets engine for a path
play_enable:
	@$(MAKE) _vault CMD="vault secrets enable -path=secret/ aws"

# configures vault to use an aws service account
play_config_aws:
	@$(MAKE) _vault CMD="vault write secret/config/root access_key=${ACCESS_KEY} secret_key=${SECRET_KEY}"

# adds an s3 crud role
play_add_s3_role:
	$(MAKE) _vault CMD='vault write secret/roles/s3-all-crud-role policy=$$(cat ./resources/s3-all-crud-policy.json)'

# retrieves a dynamic s3 crud role
play_get_s3_role:
	@$(MAKE) _vault CMD='vault read secret/creds/s3-all-crud-role'

# revoke the s3 crud role
play_revoke_s3_role:
	$(MAKE) _vault CMD="vault lease revoke -prefix secret/creds/s3-all-crud-role"

# removes the aws dynamic secrets config
play_delete_aws:
	@$(MAKE) _vault CMD="vault secrets disable secret/"

_vault:
	@docker exec go_demo_init_vault_dynamic_secrets ${CMD}