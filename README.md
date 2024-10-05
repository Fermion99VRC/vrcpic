# vrcpic <!-- omit in toc -->

<p align="center">
<a href="https://github.com/Fermion99VRC/vrcpic/actions?query=workflow%3ATest" target="_blank">
  <img src="https://github.com/Fermion99VRC/vrcpic/workflows/Test/badge.svg" alt="Test">
</a>
<a href="https://codecov.io/gh/Fermion99VRC/vrcpic" > 
  <img src="https://codecov.io/gh/Fermion99VRC/vrcpic/graph/badge.svg?token=MLEVYMQ192" alt="Coverage"> 
</a>
</p>

*vrcpic*は VRChat のスクリーンショットのためのコマンドラインツールです。

多くの VRChat プレイヤーは一日にたくさんのスクリーンショットを取ると思います。
たった数日間放置しただけでデフォルトの保存フォルダがいっぱいになったという経験もあるでしょう。
*vrcpic*はそのような問題を解決するために開発されました。

- [使い方](#使い方)
  - [スクリーンショットを分類する](#スクリーンショットを分類する)
- [License](#license)
- [Contact](#contact)

## 使い方

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

### スクリーンショットを分類する

写真を分類するには`assort`コマンドを使用します。
`-f`または`--from`フラッグで VRChat アプリがスクリーンショットを保存するデフォルトのフォルダを指定し、`-t`または`--to`フラッグで分類先のフォルダを指定してください。

```console
$ vrcpic assort -f <path/of/pictures> -t <path/to/save>
```

例えば,

```console
$ vrcpic assort -f C:\Users\hoge\Pictures\VRChat -t D:\VRChat\Pictures
```

スクリーンショットは撮られた日時によって分類されます。
午前中に撮影されたものはその日の写真として、午後に撮影されたものは次の日の写真として分類されます。
例えば、2024年4月1日の 06:00 に撮影された写真は`2024-04/0401`というフォルダに分類され、
2024年4月1日の 18:00 に撮影された写真は`2024-04/0402`というフォルダに分類されます。
これは VRChat プレイヤーが深夜0時を跨いで活動するのを考慮した仕様です。

もし元のフォルダに写真を残しておきたければ、`--keep`フラッグを追加してください。

```console
$ vrcpic assort --keep -f <path/of/pictures> -t <path/to/save>
```

詳細は`vrcpic assort --help`を実行してください。

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
