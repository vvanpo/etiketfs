# Example usages

Configures the storage location to use the `local` storage driver pointed to the `my-vind` folder in your home directory:
```console
$ vind --set-storage 'local:~/my-vind'
```
Configures a hypothetical `grpc` driver for a remote `vind` gRPC service hosted at `example.com/my-vind`:
```console
$ vind --set-storage 'grpc:example.com/my-vind'
```
The storage location is always represented by the storage type and address separated by a colon. The `--set-storage` parameter writes to a configuration file in `${XDG_HOME_CONFIG}/vind.toml`.

Without arguments, the entire catalogue of files for the configured storage is listed in a random order, displaying only `system` metadata and the `binary` format metadata. For interactive terminals the output is automatically piped to the default pager.
```console
$ vind
system/added    system/modified system/accessed system/format   binary/size binary/sha256
19 Aug 19:51    19 Aug 19:51    30 Nov 17:08    utf8,ascii      165         d2c0d9d3ca...
19 Aug 21:24    17:08:27        17:08:27        jpeg            5.8 MiB     b96b745460...
30 Nov 19:13    30 Nov 19:13    30 Nov 19:13    tiff            17 MiB      64f52d7175...
31 Jul 2018     31 Jul 2018     11 Nov 2023     gzip            161 MiB     744d1d2074...
19 Aug 14:09    19 Aug 14:09    30 Nov 17:08    utf16,json      238         4ae9812532...
19 Feb 21:07    4 Sep 21:18     17:08:54        utf8,markdown   1.1 KiB     b2ea0b0731...
```

Overrides the configured storage location:
```console
$ vind --storage 'local:/some/other/storage/path'
system/added    system/modified ...
19 Aug 19:51    19 Aug 19:51    ...
...
```

Restricts output to files matching a format:
```console
$ vind --format 'utf8'
SIZE    HASH            ...
165     d2c0d9d3ca...
238     4ae9812532...
1.1 KiB b2ea0b0731...
```

Restricts output to files matching multiple formats:
```console
$ vind --format 'utf8,markdown'
SIZE    HASH            ...
1.1 KiB b2ea0b0731...
```

Displays all unparameterized metadata from the passed formats (implies `--format`):
```console
$ vind --display 'utf8'
CHARACTERS
165
238
1073
```

```console
$ vind
```
