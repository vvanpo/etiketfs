# etiketfs
A non-hierarchical, format-aware filesystem

The filesystem exposes a collection of files, which can be queried by content
and metadata. It follows a plugin model where users can install plugins for each
file format they wish to use, and these plugins are able to decipher which files
they are responsible for and enumerate their queryable data. For example, a PDF
plugin might be able to expose the title of each document that has one, and an
MP3 plugin might be able to calculate the beats per minute for some songs. A
filesystem user with format plugins installed can query their files by filetype
and the attributes they expose.

While metadata that is inherent to (and derived directly from) file contents is
the purview of format plugins, the filesystem manages all other metadata, which
it divides into two categories: intrinsic and extrinsic. Intrinsic metadata
consists of inextricable descriptors that stay with the file as it's copied or
transmitted; e.g. the date an article was published, the timestamp when a photo
was taken, the name of a book's author, the hardware used to create a file, etc.
Extrinsic metadata is for anything personal to the user and their system's usage
of a file, which might include timestamps for when it was added to or modified
within the filesystem, access control attributes, and tags for categorization.

Unlike traditional filesystems, etiketfs does not include the concept of
location—there are no directories and files do not have filenames.

## Glossary

* *File*: A combination of content (a sequence of bytes) and metadata.

* *Filesystem*: A collection of files, and the resources needed to persist them.

  The filesystem presents an interface for browsing and filtering files by
  metadata, and by content for file formats that lend itself to querying.

* *Filter*: A predicate applied to the set of files in the filesystem to select
  a subset.

* *Format*: A description of a file's content as belonging to group of
  consistently-structured files.

  The filesystem is extended by a registry of format plugins, each of which
  describe and identify a particular format. Files in the filesystem can be
  filtered by the formats in the registry, but each format can also provide
  derived metadata values and query operations. For example, some text document
  formats might provide a `word count` metadata value, and the ability to search
  for documents by substring match.

  When a file is added to the filesystem, it is associated with the format that
  identifies it. If a file can't be identified by any of the formats in the
  registry, it will be excluded from any format-specific filters. When the file
  content is modified (which could mean it no longer satisfies its associated
  format's constraints), format re-indexing isn't guaranteed to be repeated
  immediately. If a format operation (like calculating a derived metadata value)
  fails due to format mismatch, the format association will just be removed and
  the file marked for re-indexing.

  It's possible for a file to have multiple formats, e.g. in cases where one
  format is a superset of another.

* *Metadata*: File metadata describes, identifies, or groups files.

  Metadata that is considered to belong to the file is called "intrinsic"
  metadata, and should consist of inherent attributes of the content that are
  invariant across systems.

  Moving a file into or out of a filesystem preserves intrinsic—but not
  extrinsic—metadata.

* *Storage*: The underlying interface responsible for file and metadata
  persistence.
