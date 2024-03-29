# anise-portage-converter

Portage/Overlay converter for Luet specs.
Old name `luet-portage-converter`.

```
$ anise portage-converter -- --help
Copyright (c) 2020-2023 - Daniele Rondina

Portage/Overlay converter for Luet specs.

Usage:
  anise-portage-converter -- [flags]

Flags:
      --backend string               Select backend resolver: qdepends|reposcan. (default "reposcan")
  -h, --help                         help for anise-portage-converter
      --ignore-missing-deps          Ignore missing deps on resolver.
      --override                     Override existing specs if already present.
      --reposcan-files stringArray   Append additional reposcan files. Only for reposcan backend.
      --rules string                 Rules file.
      --to string                    Targer tree where bump new specs.
  -t, --tree stringArray             Path of the tree to use.
  -v, --version                      version for anise-portage-converter

```


`anise-portage-converter` at the moment has two different backend to generate anise specs:

  * `qdepends`: based on Gentoo qdepends tool. This backend requires that the packages must be installed on the system where is executed anise-portage-converter. It doesn't work well with overlays.
     This backend will be soon set as *deprecated*.

  * `reposcan`: based on Funtoo `reposcan` generated files. It uses the JSON files generated and doesn't require that the packages must be installed :champagne: and it support Gentoo Overlay and Funtoo Kits.


## Specifications

`anise-portage-converter` take in input a YAML rules file that is used to generate anise specs.

### Example

```yaml
skipped_resolutions:
  packages:
    - name: elt-patches
      category: app-portage

    - name: eselect-opengl
      category: app-eselect

  categories:
    - virtual

#include_files:
#  - artefacts/file1.yml

build_template_file: templates/build-pkg.yaml.tmpl

artefacts:
- tree: multi-arch/packages
  packages:
    - sys-devel/gcc:9.3.0
    - app-emulation/lxd
    - app-emulation/lxc

# Define the list of the files generated by Funtoo reposcan tool.
reposcan_sources:
  - /tmp/kit_cache/gentoo-master
  - /tmp/kit_cache/geaaru_overlay-master

reposcan_contraints:
  packages:
    - dev-lang/python:3.7
    - sys-devel/gcc:9.3.0

reposcan_disabled_use_flags:
  - test
```

## Create anise specs

```
$> anise-portage-converter -- -t ./anise-packages-tree/ --to ./anise-packages-tree/ --rules config.yaml  --ignore-missing-deps

```
