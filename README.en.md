# vrcpic <!-- omit in toc -->

<p align="center">
<a href="https://github.com/Fermion99VRC/vrcpic/actions?query=workflow%3ATest" target="_blank">
  <img src="https://github.com/Fermion99VRC/vrcpic/workflows/Test/badge.svg" alt="Test">
</a>
<a href="https://codecov.io/gh/Fermion99VRC/vrcpic" > 
  <img src="https://codecov.io/gh/Fermion99VRC/vrcpic/graph/badge.svg?token=MLEVYMQ192" alt="Coverage"> 
</a>
</p>

_vrcpic_ is a command line tool for screenshots of VRChat.

Most of VRChat players may take a lot of screenshots in a day.
Without moving them from the default directory just in a few day,
the directory will be full of the screenshots.
_vrcpic_ is created to solve such a problem.

- [How to use](#how-to-use)
  - [Assort screenshots](#assort-screenshots)
- [License](#license)
- [Contact](#contact)

## How to use

```console
$ vrcpic --help
vrcpic is a CLI tool for screenshots of VRChat.

Usage:
  vrcpic [command]

Available Commands:
  assort      assort screenshots of VRChat
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command

Flags:
  -h, --help      help for vrcpic
  -v, --version   version for vrcpic

Use "vrcpic [command] --help" for more information about a command.
```

### Assort screenshots

Use `assort` command to assort screenshots.
Set your dafault directory that VRChat app saves screenshot with `-f` or `--from` flag, and your target directory to assort screenshots with `-t` or `--to` flag.

```console
$ vrcpic assort -f <path/of/screenshots> -t <path/to/save>
```

For example,

```console
$ vrcpic assort -f C:\Users\hoge\Pictures\VRChat -t D:\VRChat\Pictures
```

The screenshots are assorted by the time when they are taken.
The picture taken in the morning is assorted to the picture of the day,
and one taken in the afternoon is assorted to the picture of the next day.
This is the specification of considering some VRChat players play over the midnight.

If you want to keep the pictures in the original directory, add `--keep` flag.

```console
$ vrcpic assort --keep -f <path/of/pictures> -t <path/to/save>
```

For more information, run `vrcpic assort --help`.

```console
$ vrcpic assort --help
Move screenshots of VRChat to <to>/YYYY-MM/MMDD.
<to> is given with the flag "--to".
Note that the screenshots taken after 12:00 are moved to the next-day directory.
For example, the screenshot taken at 18:00 on 2024-01-01 will be moved to <to>/2024-01/0102.

Usage:
  vrcpic assort [flags]

Flags:
  -f, --from string   directory where VRChat app saves screenshots (default ".")
  -h, --help          help for assort
      --keep          keep screenshots in the original directory
  -t, --to string     directory to which screenshots are moved (default ".")
```

## License

[MIT License](./LICENSE)

## Contact

X(Twitter): [@F99vrc](https://x.com/F99vrc)
