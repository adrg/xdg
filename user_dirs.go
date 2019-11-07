package xdg

// XDG user directories environment variables.
var (
	envDesktopDir     = "XDG_DESKTOP_DIR"
	envDownloadDir    = "XDG_DOWNLOAD_DIR"
	envTemplatesDir   = "XDG_TEMPLATES_DIR"
	envPublicShareDir = "XDG_PUBLICSHARE_DIR"
	envDocumentsDir   = "XDG_DOCUMENTS_DIR"
	envMusicDir       = "XDG_MUSIC_DIR"
	envPicturesDir    = "XDG_PICTURES_DIR"
	envVideosDir      = "XDG_VIDEOS_DIR"
)

// UserDirectories defines the locations of well known user directories.
type UserDirectories struct {
	// Desktop defines the location of the user's desktop directory.
	Desktop string

	// Download defines a suitable location for user downloaded files.
	Download string

	// Templates defines a suitable location for user template files.
	Templates string

	// PublicShare defines a suitable location for user shared files.
	PublicShare string

	// Documents defines a suitable location for user document files.
	Documents string

	// Music defines a suitable location for user audio files.
	Music string

	// Pictures defines a suitable location for user image files.
	Pictures string

	// VideosDir defines a suitable location for user video files.
	Videos string
}

// DesktopFile returns a suitable location for the specified file on the
// user's desktop. The relPath parameter must contain the name of the file,
// and optionally, a set of parent directories (e.g. dir/file.ext).
// If the specified directories do not exist, they will be created relative
// to the user's desktop directory. On failure, an error containing the
// attempted paths is returned.
func (ud UserDirectories) DesktopFile(relPath string) (string, error) {
	return createPath(relPath, []string{ud.Desktop})
}

// DownloadedFile returns a suitable location for the specified file in the
// user's download directory. The relPath parameter must contain the name of
// the file, and optionally, a set of parent directories (e.g. dir/file.ext).
// If the specified directories do not exist, they will be created relative
// to the user's download directory. On failure, an error containing the
// attempted paths is returned.
func (ud UserDirectories) DownloadedFile(relPath string) (string, error) {
	return createPath(relPath, []string{ud.Download})
}

// TemplateFile returns a suitable location for the specified file in the
// user's templates directory. The relPath parameter must contain the name of
// the file, and optionally, a set of parent directories (e.g. dir/file.tpl).
// If the specified directories do not exist, they will be created relative
// to the user's templates directory. On failure, an error containing the
// attempted paths is returned.
func (ud UserDirectories) TemplateFile(relPath string) (string, error) {
	return createPath(relPath, []string{ud.Templates})
}

// PublicFile returns a suitable location for the specified file in the
// user's public directory. The relPath parameter must contain the name of
// the file, and optionally, a set of parent directories (e.g. dir/file.ext).
// If the specified directories do not exist, they will be created relative
// to the user's public directory. On failure, an error containing the
// attempted paths is returned.
func (ud UserDirectories) PublicFile(relPath string) (string, error) {
	return createPath(relPath, []string{ud.PublicShare})
}

// DocumentFile returns a suitable location for the specified file in the
// user's documents directory. The relPath parameter must contain the name of
// the file, and optionally, a set of parent directories (e.g. dir/file.txt).
// If the specified directories do not exist, they will be created relative
// to the user's documents directory. On failure, an error containing the
// attempted paths is returned.
func (ud UserDirectories) DocumentFile(relPath string) (string, error) {
	return createPath(relPath, []string{ud.Documents})
}

// AudioFile returns a suitable location for the specified file in the
// user's music directory. The relPath parameter must contain the name of
// the file, and optionally, a set of parent directories (e.g. dir/file.mp3).
// If the specified directories do not exist, they will be created relative
// to the user's music directory. On failure, an error containing the
// attempted paths is returned.
func (ud UserDirectories) AudioFile(relPath string) (string, error) {
	return createPath(relPath, []string{ud.Music})
}

// ImageFile returns a suitable location for the specified file in the
// user's pictures directory. The relPath parameter must contain the name of
// the file, and optionally, a set of parent directories (e.g. dir/file.png).
// If the specified directories do not exist, they will be created relative
// to the user's pictures directory. On failure, an error containing the
// attempted paths is returned.
func (ud UserDirectories) ImageFile(relPath string) (string, error) {
	return createPath(relPath, []string{ud.Pictures})
}

// VideoFile returns a suitable location for the specified file in the
// user's videos directory. The relPath parameter must contain the name of
// the file, and optionally, a set of parent directories (e.g. dir/file.mkv).
// If the specified directories do not exist, they will be created relative
// to the user's videos directory. On failure, an error containing the
// attempted paths is returned.
func (ud UserDirectories) VideoFile(relPath string) (string, error) {
	return createPath(relPath, []string{ud.Videos})
}

// SearchDesktopFile searches for the specified file on the user's desktop
// directory. The relPath parameter must contain the name of the file, and
// optionally, a set of parent directories (e.g. dir/file.ext). If the
// file cannot be found, an error specifying the searched path is returned.
func (ud UserDirectories) SearchDesktopFile(relPath string) (string, error) {
	return searchFile(relPath, []string{ud.Desktop})
}

// SearchDownloadedFile searches for the specified file in the user's download
// directory. The relPath parameter must contain the name of the file, and
// optionally, a set of parent directories (e.g. dir/file.ext). If the
// file cannot be found, an error specifying the searched path is returned.
func (ud UserDirectories) SearchDownloadedFile(relPath string) (string, error) {
	return searchFile(relPath, []string{ud.Download})
}

// SearchTemplateFile searches for the specified file in the user's templates
// directory. The relPath parameter must contain the name of the template file,
// and optionally, a set of parent directories (e.g. dir/file.tpl). If the
// file cannot be found, an error specifying the searched path is returned.
func (ud UserDirectories) SearchTemplateFile(relPath string) (string, error) {
	return searchFile(relPath, []string{ud.Templates})
}

// SearchPublicFile searches for the specified file in the user's public
// directory. The relPath parameter must contain the name of the file, and
// optionally, a set of parent directories (e.g. dir/file.ext). If the
// file cannot be found, an error specifying the searched path is returned.
func (ud UserDirectories) SearchPublicFile(relPath string) (string, error) {
	return searchFile(relPath, []string{ud.PublicShare})
}

// SearchDocumentFile searches for the specified file in the user's documents
// directory. The relPath parameter must contain the name of the document file,
// and optionally, a set of parent directories (e.g. dir/file.txt). If the
// file cannot be found, an error specifying the searched path is returned.
func (ud UserDirectories) SearchDocumentFile(relPath string) (string, error) {
	return searchFile(relPath, []string{ud.Documents})
}

// SearchAudioFile searches for the specified file in the user's music
// directory. The relPath parameter must contain the name of the audio file,
// and optionally, a set of parent directories (e.g. dir/file.mp3). If the
// file cannot be found, an error specifying the searched path is returned.
func (ud UserDirectories) SearchAudioFile(relPath string) (string, error) {
	return searchFile(relPath, []string{ud.Music})
}

// SearchImageFile searches for the specified file in the user's pictures
// directory. The relPath parameter must contain the name of the image file,
// and optionally, a set of parent directories (e.g. dir/file.png). If the
// file cannot be found, an error specifying the searched path is returned.
func (ud UserDirectories) SearchImageFile(relPath string) (string, error) {
	return searchFile(relPath, []string{ud.Pictures})
}

// SearchVideoFile searches for the specified file in the user's videos
// directory. The relPath parameter must contain the name of the video file,
// and optionally, a set of parent directories (e.g. dir/file.mkv). If the
// file cannot be found, an error specifying the searched path is returned.
func (ud UserDirectories) SearchVideoFile(relPath string) (string, error) {
	return searchFile(relPath, []string{ud.Videos})
}
