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
