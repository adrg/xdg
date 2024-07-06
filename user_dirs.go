package xdg

// UserDirectories defines the locations of well known user directories.
type UserDirectories struct {
	// Desktop defines the location of the user's desktop directory.
	Desktop string

	// Download defines a suitable location for user downloaded files.
	Download string

	// Documents defines a suitable location for user document files.
	Documents string

	// Music defines a suitable location for user audio files.
	Music string

	// Pictures defines a suitable location for user image files.
	Pictures string

	// VideosDir defines a suitable location for user video files.
	Videos string

	// Templates defines a suitable location for user template files.
	Templates string

	// PublicShare defines a suitable location for user shared files.
	PublicShare string
}
