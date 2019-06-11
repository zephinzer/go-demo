include ../Makefile

# sets up the environment
play:
	@docker-compose down
	@docker-compose up -d -V

# tears down the environment
play_stop:
	@docker-compose down

# use this to initialise vault
play_init:
	@docker exec -it go_demo_init_vault vault operator init \
		-key-shares=3 \
		-key-threshold=2
	# copy the Unseal Keys and Root Token and paste it somewhere!

# use this to check the status of vault
play_status:
	@docker exec -it go_demo_init_vault vault status

# use this to unseal vault (use keys from `make play_init`)
play_unseal:
	# run this 3 times using a different Unseal Key from `play_init` each time
	@docker exec -it go_demo_init_vault vault operator unseal

# use this to login to vault via the cli (use root token from `make play_init`)
play_login:
	# enter the Root Token from `play_init` here
	@docker exec -it go_demo_init_vault vault login

# use this to enable the key-value (kv) secrets engine for a path
play_enable:
	@docker exec -it go_demo_init_vault vault secrets enable -path=secret/ kv

# use this to list secrets
play_list:
	@docker exec -it go_demo_init_vault vault list secret/

# demonstrates adding a new secret
play_put:
	@docker exec -it go_demo_init_vault vault kv put secret/s hello=world

# demonstrates retrieval of a secret
play_get:
	@docker exec -it go_demo_init_vault vault kv get secret/s

# demonstrates retrieval of a secret in json format
play_get_json:
	@docker exec -it go_demo_init_vault vault kv get -format=json secret/s

# retrieves a specified version (V) of a secret
play_get_version:
	@docker exec -it go_demo_init_vault vault kv get -version=${V} secret/s
	# if you got an error, it's likely V is not defined, use `make play_get_version V=x` where x is a number
	# from 0 to the max version provided by `make play_get`

# demonstrates enabling of versioning of a secret
play_enable_versioning:
	@docker exec -it go_demo_init_vault vault kv enable-versioning secret

# demonstrates updating the secret we created (sets it to spanish)
play_patch_spanish:
	@docker exec -it go_demo_init_vault vault kv put secret/s hola=espanol

# demonstrates updating the secret we created (sets it to german)
play_patch_german:
	@docker exec -it go_demo_init_vault vault kv put secret/s halo=deutcsh

# removes a secret
play_delete:
	@docker exec -it go_demo_init_vault vault kv delete secret/s
