# codename-generator

codename-generator is a CLI tool and Golang library that generates unique codenames from text.

## Installation

```
go install github.com/ToshihitoKon/codename-generator@latest
```

## Usage

```
 % codename-generator --help
 Usage of codename-generator:
   -g, --generator-version int   Generator Version (default -1)
   -h, --help                    show help
   -l, --list                    List generator versions
```

`generator-version` option is required.

```
% echo codename-generator-v1.0 | codename-generator --generator-version 1
Astrophysics Miranda

 % codename-generator --generator-version 1 "codename-generator-v1.0"
Astrophysics Miranda
```
