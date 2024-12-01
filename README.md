# etiketfs
A non-hierarchical, format-aware filesystem

The filesystem exposes a collection of files, which can be queried by content
and metadata. It follows a plugin model where users can install plugins for each
file format they wish to use, and these plugins are able to decipher which files
they are responsible for and enumerate their queryable data. For example, a PDF
plugin might be able to expose the title of each document that has one, and an
MP3 plugin might be able to calculate the beats per minute for some songs. A
filesystem user with format plugins installed can query their files by format
and the properties they expose.

While metadata that is inherent to (and derived directly from) file content is
the purview of format plugins, the filesystem manages all other metadata. Users
can categorize their files with labels, add or update descriptive fields, and
describe relationships to other files. Users can compose filter and sort
operations using this metadata to produce file selections for browsing.

Unlike traditional filesystems, etiketfs does not include the concept of
location—there are no directories and files do not have filenames. This presents
an acute limitation in that it is not compatible with operating system
filesystem APIs, meaning applications must be specifically designed to interface
with etiketfs.

## Glossary

* *File*: A combination of **content** (a sequence of bytes) and **metadata**.

* *Filesystem*: A collection of files and the resources needed to persist them.

* *Filter*: An operation mapping a selection to a subselection, via a predicate
  applied on metadata properties.

* *Format*: A description of a file's content as belonging to a group of
  consistently-structured files.

  The filesystem is extended by a set of format plugins, each of which describe
  and identify a particular file format. A format plugin enumerates derived
  metadata properties for a file and is responsible for parsing or calculating
  their values. When a file is written to or added to the filesystem, it is
  associated with formats that identify it.

  To prevent collisions when a file is associated with more then one format,
  metadata identifiers provided by formats are namespaced using the format name.

  A format can be super/subset, as in cases where one format is
  backwards-compatible with another (e.g. UTF-8 and ASCII, YAML and JSON).
  Formats that extend another only have their identifying function called on
  files that have already been matched by all their subset formats. All files
  match the `binary` format, meaning all other formats extend `binary`.

  **Format groups** are named groups of format implementations that present one
  or more (semantically) identical metadata properties. These properties are
  then instead available under the group namespace. For example, the `jpeg` and
  `png` formats are members of the `image` format group, and the `resolution`
  property determined by each format implementation is queried as
  `image/resolution`. Another example is the various Unicode formats (`utf8`,
  `utf16`, etc.) all exposing the number of characters in a document as the
  `unicode/characters` property.

* *Metadata*: File metadata describes and identifies the content of a file. A
  metadata **property** consists of an association between an **identifier** and
  a file. A property can be dereferenced to produce a metadata **value**.

  Metadata that is considered to belong to the file is called **intrinsic**
  metadata, and should consist of inherent properties of the file content that
  are invariant across systems. **Extrinsic** metadata is scoped to the
  filesystem itself, representing something that is only useful within the
  context of the system and the other files in it. Moving a file into or out of
  a filesystem preserves intrinsic—but not extrinsic—metadata.

  **Derived** metadata consists of read-only properties calculated from the file
  content, and are therefore intrinsic. Some derived properties may accept a
  user-supplied argument, whose value is used in the property calculation.
  Parameterized property identifiers are grouped by name and the defined
  parameter, where each argument value produces a unique identifier.

  **Attributive** metadata consists of properties (attributes) that describe or
  relate to the file but are otherwise orthogonal to its content; they are
  "attributed" to the file content by someone or something. Extrinsic attributes
  are stateful and managed by the filesystem, while intrinsic attributes are
  exclusive to files that support an in-band metadata container format such as
  Exif or ID3—intrinsic attributes are thus exposed via a corresponding format
  plugin.

  Properties originate from a number of sources:
  * The filesystem provides a handful of universal read-only properties, like
    added/modified/accessed timestamps.
  * Each format can expose a variable number of intrinsic properties for files
    they identify.
  * The user can define any number of mutable extrinsic attributes.

  Metadata values and derived property parameters are typed, determining how
  they are represented and how they can be operated on.

* *Selection*: An unordered subset of files in the filesystem.

* *Sort*: A compare operation applied to a selection, accepting a metadata
  property common to files to return an ordered list of subselections.

* *Storage*: The underlying interface responsible for file and metadata
  persistence.

## To-dos

* Reactive updates in the selection and sorted selection interfaces triggered by
  state changes. This can be caused by changes in the input selection (i.e. a
  file that is in the output selection is removed from the input selection, or a
  file is added to the input selection that would not be filtered out), or a
  change to the metadata values filtered on (i.e. a file in the input selection
  that is filtered out updates to be included, or a file that is included no
  longer should be).

* Implement the ability to make intrinsic attributes modifiable through the
  filesystem interface.

* Design "format-unwrapping" to expose metadata from formats that are
  encapsulated within another. In particular, this would allow compressed files
  to expose the same metadata as when uncompressed. Perhaps format
  implementations for compression formats (gzip, bzip2, lzma, etc.) could
  toggle a `compressed` system metadata property and provide a decompression
  interface to allow other formats to parse the uncompressed content.
