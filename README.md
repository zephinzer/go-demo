# Demos
A playground for container orchestration and cloud native applications.

[![Build Status](https://travis-ci.org/zephinzer/go-demo.svg?branch=master)](https://travis-ci.org/zephinzer/go-demo)

# Usage

The code is mostly self-documenting via the `Makefile`s in each directory. The convention in each directory enables you to run `make` in any of the non-WIP directories and see a help text output. Open the relevant `Makefile`s to check out what can be done.

If something's unclear/not right, feel free to [raise an issue](/zephinzer/go-demo/issues)!

# Content

- [Example Deployments](./deployments/README.md)
- [Cluster Initialisations](./init/README.md)
- [Cloud Native Demo Tools](./tools/README.md)

## Setting Up

Run `make` to create all necessary binaries and Docker images.

Run `make ssl` to create the required certificates/keys to support HTTPS.

Run `make showcase` to create the Docker Compose setup.

# License
Feel free to use these examples in your own workshops/tutorials!

Code and usage of its resulting binaries is licensed under the permissive [MIT license](./LICENSE)

Content is licensed under the [Creative Commons Attribution-ShareAlike 3.0 (CC BY-SA 3.0 SG) license](https://creativecommons.org/licenses/by-sa/3.0/sg/). For attribution, use the following:

## HTML

```html
Content was created by <a href="https://github.com/zephinzer" target="_blank">Joseph Matthias Goh/@zephinzer</a> and the original content can be found at <a href="https://github.com/zephinzer/go-demo" target="_blank">https://github.com/zephinzer/go-demo</a>.
```

## Markdown

```markdown
Content was created by [Joseph Matthias Goh/@zephinzer](https://github.com/zephinzer) and the original content can be found at [https://github.com/zephinzer/go-demo](https://github.com/zephinzer/go-demo).
```

# Cheers!
