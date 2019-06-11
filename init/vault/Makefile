include ../Makefile

play:
	@docker-compose down
	@docker-compose up -d -V

play_stop:
	@docker-compose down

play_init:
	@docker exec -it go_demo_init_vault vault operator init
	# copy the Unseal Keys and Root Token and paste it somewhere!

play_status:
	@docker exec -it go_demo_init_vault vault status

play_unseal:
	# run this 3 times using a different Unseal Key from `play_init` each time
	@docker exec -it go_demo_init_vault vault operator unseal

play_login:
	# enter the Root Token from `play_init` here
	@docker exec -it go_demo_init_vault vault login

play_enable:
	@docker exec -it go_demo_init_vault vault secrets enable -path=secret/ kv

play_list:
	@docker exec -it go_demo_init_vault vault list secret/

play_put:
	@docker exec -it go_demo_init_vault vault kv put secret/s hello=world

play_get:
	@docker exec -it go_demo_init_vault vault kv get secret/s

play_get_json:
	@docker exec -it go_demo_init_vault vault kv get -format=json secret/s

play_get_version:
	@docker exec -it go_demo_init_vault vault kv get -version=${V} secret/s
	# if you got an error, it's likely V is not defined, use `make play_get_version V=x` where x is a number
	# from 0 to the max version provided by `make play_get`

play_enable_versioning:
	@docker exec -it go_demo_init_vault vault kv enable-versioning secret

play_patch_spanish:
	@docker exec -it go_demo_init_vault vault kv put secret/s hola=espanol

play_patch_german:
	@docker exec -it go_demo_init_vault vault kv put secret/s halo=deutcsh

play_delete:
	@docker exec -it go_demo_init_vault vault kv delete secret/s
