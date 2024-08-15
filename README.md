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

* *File*: A combination of content (a sequence of bytes) and metadata.

* *Filesystem*: A collection of files and the resources needed to persist them.

* *Filter*: An operation mapping a selection to a subselection, via a predicate
  applied on metadata properties.

* *Format*: A description of a file's content as belonging to a group of
  consistently-structured files.

  The filesystem is extended by a registry of format plugins, each of which
  describe and identify a particular format. A format provides intrinsic
  metadata properties for a file and is responsible for calculating derived
  values. When a file is added to the filesystem, it is associated with formats
  that identify it.

  To prevent collisions when a file is associated with more then one format,
  metadata identifiers provided by formats are namespaced using the format name.

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
  content. Derived property identifiers can be grouped by name and a defined
  parameter, where each argument value produces a unique identifier. The
  argument value is then used in derived property calculation.

  **Attributive** metadata consists of mutable properties that describe or
  relate to the file but are otherwise orthogonal to its content; they are
  managed by the filesystem and are not affected by changes to file content.

  Metadata values and derived property parameters are typed, using a handful of
  data types.

  Properties originate from a number of sources:
  * The filesystem provides a number of universal properties, like the list of
    identified formats, file size, and added/modified/accessed timestamps.
  * Each format enumerates the set of derived properties it can calculate on
    files.
  * The user can create any number of attributive properties on a file.
  * Format importers strip encoded attributive metadata (e.g. Exif) from the
    file content as a file is added/written to, and set them as modifiable
    attributes.

  Each property source scopes the property names they define, preventing
  collisions.

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
  that is filtered out updates to be included, or a file that is include no
  longer should be).

* Provide support for writing to open files. Format identification and property
  parsing/calculation would need to be done lazily, as it is an I/O and
  computationally expensive process.
