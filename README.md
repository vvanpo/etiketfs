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

* *Filesystem*: A collection of files, metadata, and the resources needed to
  persist them. The filesystem presents an interface for browsing, adding or
  removing, and manipulating the extrinsic metadata of files.

* *Filter*: A predicate applied to the set of files in the filesystem to produce
  a subset. Multiple filters can be composed to generate a selection.

  Filters operate on a selection of files and a metadata property common to all
  files in the selection, using the values of the property to filter out some
  files and produce a subselection. The first filter in a chain must operate on
  the seletion of all files in the filesystem.

* *Format*: A description of a file's content as belonging to a group of
  consistently-structured files.

  The filesystem is extended by a registry of format plugins, each of which
  describe and identify a particular format. Files in the filesystem can be
  filtered by the formats in the registry, but each format can also provide
  derived metadata values. For example, some text document formats might provide
  a word count metadata property, and a text match property accepting a string
  parameter.

  When a file is added to the filesystem, it is associated with the format that
  identifies it. If a file can't be identified by any of the formats in the
  registry, it will be excluded from any format-specific filters. When the file
  content is modified (which could mean it no longer satisfies its associated
  format's constraints), format re-indexing isn't guaranteed to be repeated
  immediately. If a format operation (like calculating a derived metadata value)
  fails due to format mismatch, the format association will be removed and the
  file marked for re-indexing.

  It's possible for a file to have multiple formats, e.g. when one format is a
  superset of another.

* *Metadata*: File metadata describes and identifies the content of a file. A
  metadata property consists of an **identifier** and a **value**.

  **Attributive** metadata consists of stateful properties that users can add,
  remove, or modify.

  **Derived** metadata consists of read-only properties calculated from the file
  content, and any other arguments the property might take. The identifier of a
  derived property consists of both name and arguments.

  Metadata that is considered to belong to the file is called **intrinsic**
  metadata, and should consist of inherent properties of the file content that
  are invariant across systems.

  **Extrinsic** metadata is scoped to the filesystem itself, representing
  something that is only useful within the context of the system and the other
  files in it. For example a file could be given a unique label for ease of
  filtering, assigned membership to one or more categories, or point to another
  file to establish a relationship.

  Moving a file into or out of a filesystem preserves intrinsic—but not
  extrinsic—metadata.

  Metadata values are typed, using a handful of scalar data types. The type
  determines the availability of filtering and sorting operations for a metadata
  property. For example, a date-typed property might be filterable using date
  ranges, and sortable in ascending or descending order.

* *Selection*: An immutable unordered subset of files in the filesystem.
  Selections are the result of a filter operation on another selection, with the
  originating selection in a chain of filters being the complete set of files in
  the filesystem.

  Selections are retrieved from the filesystem and filters via channels, which
  generate selections reactively. When there is a state change in the filesystem
  that would alter the requested selection (e.g. a file in the selection is
  removed, a metadata property being filtered on changes in value so as to be
  excluded from the filter, etc.), an updated selection is sent to the channel.

  A sort operation can be applied to a selection to return an ordered list of
  files.

* *Storage*: The underlying interface responsible for file and metadata
  persistence.
