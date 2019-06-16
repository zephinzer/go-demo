> [Go-Demo](../../../) > [Deployments](../../) > [WordPress](../) > Docker Compose

# Wordpress (Docker Compose)

This directory contains a Docker Compose that brings up a functiona Wordpress installation.

| Service | Accessible At |
| --- | --- |
| `wordpress` | http://localhost:18000 |
| `mysql` | http://localhost:13306 |

# Run-Through

To start it, use:

```sh
docker-compose up -V
```

Open a browser and navigate to [http://localhost:18080](http://localhost:18080) and you should see the WordPress installation page. Fill it up with the following details (or change it, up to you):

| Field | Value |
| --- | --- |
| **Site Title** | Docker deployed WordPress |
| **Username** | username |
| **Password** | password |
| **Confirm Password** | `true` (don't use the password "password" in production obviously) |
| **Your Email** | test@domain.com |
| **Search Engine Visibility** | `true` |

You should be brought to the login page now. Log in using `username`:`password` and you're good to go!

# Extending: Local Theme/Plugin Development

For development of themes/plugins, mount an additional volume to your local drive for persistence. First create a directory named `wp-content`:

```sh
mkdir -p ./wp-content
```

Then add the following YAML into the Docker Compose file:

```yaml
# ...
wordpress:
  # ...
  volumes:
    # ...
    - ./wp-content:/var/www/html/wp-content
# ...
```

> Note that the contents of the wp-content directory will be created under `root` which the WordPress installation is running with. To delete/modify the contents, there's a script in the [`Makefile`](./Makefile) named `update_permissions` which you can use to reset the permissions to yours. Simply run `make update_permissions` from this directory.
