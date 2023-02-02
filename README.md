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

While metadata that is inherent to (and derived directly from) file contents is
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

* *File*: A combination of content (a sequence of bytes) and intrinsic metadata.

  Files can enumerate all their intrinsic metadata property names. The file
  format constrains which properties are available on the file, both attributive
  (writeable) and derived.

* *Filesystem*: A collection of files, metadata, and the resources needed to
  persist them.

  The filesystem can be queried for:
  * A selection of all files.
  * To enumerate the metadata properties in common amongst the files in a
    selection.
  * Change notifications for metadata properties belonging to the files in a
    selection. If the property is added, removed, or updated for a file in the
    selection, a notification will be sent.
  * Extrinsic metadata associations by identifier (i.e. list all files with
    property \<name\>) or selection (list all properties associated with files
    in \<selection\>).

* *Filter*: An operation on a selection, via a predicate applied on a metadata
  property common to them, producing a subselection. Filters are typed and can
  only operate on properties of matching type.

  Filters update reactively with filesystem state changes. This can be caused by
  changes in the input selection (i.e. a file that is in the output selection is
  removed from the input selection, or a file is added to the input selection
  that would not be filtered out), or a change to the metadata values filtered
  on (i.e. a file in the input selection that is filtered out updates to be
  included, or a file that is include no longer should be).

* *Format*: A description of a file's content as belonging to a group of
  consistently-structured files.

  The filesystem is extended by a registry of format plugins, each of which
  describe and identify a particular format. Files in the filesystem can be
  filtered by the formats in the registry, but each format can also provide
  derived metadata values.

  When a file is added to the filesystem, it is associated with the format that
  identifies it. If a file can't be identified by any of the formats in the
  registry, it will be excluded from any format-specific filters. When the file
  content is modified (which could mean it no longer satisfies its associated
  format's constraints), format re-indexing isn't guaranteed to be repeated
  immediately. If a format operation (like calculating a derived metadata value)
  fails due to format mismatch, the format association will be removed and the
  file marked for re-indexing.

  It's possible for a file to have multiple formats, e.g. when one format is a
  superset of another. When metadata name collisions occur, the order of the
  format plugins in the registry is consulted to determine precedence.

* *Metadata*: File metadata describes and identifies the content of a file. A
  metadata property consists of an **identifier** and a **value**.

  **Attributive** metadata consists of stateful properties that users can add,
  remove, or modify.

  **Derived** metadata consists of read-only properties calculated from the file
  content, and any other arguments the property might take. The identifier of a
  derived property consists of both name and arguments.

  Metadata values (and derived property arguments) are typed, using a handful of
  scalar data types.

  Metadata that is considered to belong to the file is called **intrinsic**
  metadata, and should consist of inherent properties of the file content that
  are invariant across systems. **Extrinsic** metadata is scoped to the
  filesystem itself, representing something that is only useful within the
  context of the system and the other files in it. Moving a file into or out of
  a filesystem preserves intrinsic—but not extrinsic—metadata.

* *Selection*: An immutable unordered subset of files in the filesystem.

* *Sort*: An operation applied to a selection, accepting a metadata property
  common to files to return an ordered list of files. Similar to filters, sort
  operations are typed and updated reactively.

* *Storage*: The underlying interface responsible for file and metadata
  persistence.
