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

  The filesystem is extended by a set of format plugins, each of which describe
  and identify a particular file format. A format plugin enumerates derived
  metadata properties for a file and is responsible for parsing or calculating
  their values. When a file is written to or added to the filesystem, it is
  associated with formats that identify it.

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
  argument value is then used in derived property calculation. Derived
  properties are always intrinsic.

  **Attributive** metadata consists of mutable properties that describe or
  relate to the file but are otherwise orthogonal to its content; they are not
  affected by changes to file content.

  The distinction between derived and attributive metadata breaks down when
  considering in-band metadata container formats like Exif and ID3, as their
  purpose is to persist intrinsic attributes as part of the file content.
  Because they are part of the content, a format plugin corresponding to the
  container format is responsible for extracting them as a derived property, and
  they are not exposed as mutable attributes.

  Properties originate from a number of sources:
  * The filesystem provides a number of universal properties. Some are derived,
    like the list of identified formats, file size, or a content hash, and some
    are attributive, like added/modified/accessed timestamps.
  * Each format enumerates the set of properties it can derive from a file.
  * The user can define any number of extrinsic attributes.

  Each property source scopes the property names they define, preventing
  collisions.

  Metadata values and derived property parameters are typed, using a handful of
  data types.

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
