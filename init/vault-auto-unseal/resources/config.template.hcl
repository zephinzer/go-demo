#
# WARNING: DO NOT COMMIT FILE WITH SECRETS
#
# for this to work, replace the following:
#
# __region__ : aws region
# __access_key__ : aws iam access key
# __secret_key__ : aws iam secret key
# __kms_key_id__ : aws kms id

backend "file" {
  path = "/vault/file"
}

listener "tcp" {
  address = "0.0.0.0:8200"
  tls_disable = 1
}

seal "awskms" {
  region = __region__
  access_key = __access_key__
  secret_key = __secret_key__
  kms_key_id = __kms_key_id__
}

default_lease_ttl = "168h"
log_level = "trace"
max_lease_ttl = "720h"
ui = true
