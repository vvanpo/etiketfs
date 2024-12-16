# Example usages

Configures the storage location to use the `local` storage driver pointed to the `my-vind` folder in your home directory:
```console
$ vind --set-config 'storage=local:~/my-vind'
```
Configures a hypothetical `grpc` driver for a remote `vind` gRPC service hosted at `example.com/my-vind`:
```console
$ vind --set-config 'storage=grpc:example.com/my-vind'
```
The storage location is always represented by the storage type and address separated by a colon. The `--set-config` parameter writes to a configuration file at `${XDG_HOME_CONFIG}/vind.toml`, or to a path in the `VIND_CONFIG` environment variable if present.

Without arguments, the entire catalogue of files for the configured storage is listed in a random order, by default displaying only system metadata and the `binary/size` property. For interactive terminals the output is automatically piped to the default pager.
```console
$ vind
formats         added           modified        accessed        binary/size
utf8,ascii      19 Aug 19:51    19 Aug 19:51    30 Nov 17:08    165 B
jpeg            19 Aug 21:24    17:08:27        17:08:27        5.8 MiB
tiff            30 Nov 19:13    30 Nov 19:13    30 Nov 19:13    17 MiB
gzip            31 Jul 2018     31 Jul 2018     11 Nov 2023     161 MiB
utf16,json      19 Aug 14:09    19 Aug 14:09    30 Nov 17:08    238 B
utf8,markdown   19 Feb 21:07    4 Sep 21:18     17:08:54        1.1 KiB
```

Overrides the configured storage location:
```console
$ vind --storage 'local:/some/other/storage/path'
formats  added          ...
pdf      1 Dec 16:09    ...
...
```

Enumerates the list of available properties and their types:
```console
$ vind --properties
group       name        value type              parameter type      description
            added       timestamp                                   When the file was first added to the filesystem.
            modified    timestamp                                   When the file content was last modified.
            accessed    timestamp                                   When the file was last opened.
            format                                                  The format that best matches the file content.
            formats     set of text                                 All formats identified that can describe the file.
binary      size        bytes                                       The size of the file content.
publication title       text
publication authors     list of text                                The authors of the publication, sometimes ordered by relative contribution.
publication published   date
publication publisher   text
text        characters  non-negative integer
text        matches     boolean                 regular expression
text        words       non-negative integer
unicode     characters  non-negative integer
```

Restricts output to files matching a format:
```console
$ vind --filter 'utf8 in formats'
formats         added           ...
utf8,ascii      19 Aug 19:51    ...
utf8,markdown   19 Feb 21:07    ...
```

Restricts output to files matching multiple formats:
```console
$ vind --filter 'utf8,markdown in formats'
formats         added           ...
utf8,markdown   19 Feb 21:07    ...
```

Restricts output to files matching at least one of the specified formats:
```console
$ vind --filter 'jpeg|tiff in formats'
formats  added          ...
jpeg     19 Aug 21:24   ...
tiff     30 Nov 19:13   ...
```

Displays all unparameterized metadata from the passed formats and format groups:
```console
$ vind --display 'unicode'
unicode/characters
165
119
1073
```
Files that do not match any of the specified formats or otherwise do not contain any of their properties are elided from the output.

Displays the specified metadata properties:
```console
$ vind --display '/formats,unicode/characters'
formats         unicode/characters
utf8,ascii      165
utf16,json      119
utf8,markdown   1073
```
System properties are anchored with `/`. Files that don't contain any of the specified properties are elided from the output.

Configures the default displayed properties for when `vind` is called without the `--display` parameter:
```console
$ vind --set-config 'display=/formats,/modified,binary/size'
$ vind
formats         modified        binary/size
utf8,ascii      19 Aug 19:51    165 B
...
```

Configures the output format:
```console
$ vind --out 'json'
[
  {
    "formats": ["utf8", "ascii"],
    "modified": "19 Aug 19:51",
    "binary/size": 165
  },
...
```
The representation of properties is matched as close as possible within the output format's constraints. In this case, the `bytes` datatype is represented as a number without a unit.
