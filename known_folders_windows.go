package xdg

import (
	"path/filepath"
	"strings"

	"github.com/adrg/xdg/internal/util"
	"golang.org/x/sys/windows"
)

type knownFolders struct {
	systemDrive    string
	systemRoot     string
	programData    string
	roamingAppData string
	localAppData   string
	desktop        string
	downloads      string
	documents      string
	music          string
	pictures       string
	videos         string
	templates      string
	public         string
	fonts          string
	programs       string
	commonPrograms string
}

func initKnownFolders(home string) *knownFolders {
	sep := string(filepath.Separator)

	kf := &knownFolders{}
	kf.systemDrive = strings.TrimRight(util.KnownFolderPath(
		nil,
		[]string{"SystemDrive"},
		[]string{filepath.VolumeName(home), `C:\`},
	), sep) + sep
	kf.systemRoot = util.KnownFolderPath(
		windows.FOLDERID_Windows,
		[]string{"SystemRoot", "windir"},
		[]string{filepath.Join(kf.systemDrive, "Windows")},
	)
	kf.programData = util.KnownFolderPath(
		windows.FOLDERID_ProgramData,
		[]string{"ALLUSERSPROFILE", "ProgramData"},
		[]string{filepath.Join(kf.systemDrive, "ProgramData")},
	)
	kf.roamingAppData = util.KnownFolderPath(
		windows.FOLDERID_RoamingAppData,
		[]string{"APPDATA"},
		[]string{filepath.Join(home, "AppData", "Roaming")},
	)
	kf.localAppData = util.KnownFolderPath(
		windows.FOLDERID_LocalAppData,
		[]string{"LOCALAPPDATA"},
		[]string{filepath.Join(home, "AppData", "Local")},
	)
	kf.desktop = util.KnownFolderPath(
		windows.FOLDERID_Desktop,
		nil,
		[]string{filepath.Join(home, "Desktop")},
	)
	kf.downloads = util.KnownFolderPath(
		windows.FOLDERID_Downloads,
		nil,
		[]string{filepath.Join(home, "Downloads")},
	)
	kf.documents = util.KnownFolderPath(
		windows.FOLDERID_Documents,
		nil,
		[]string{filepath.Join(home, "Documents")},
	)
	kf.music = util.KnownFolderPath(
		windows.FOLDERID_Music,
		nil,
		[]string{filepath.Join(home, "Music")},
	)
	kf.pictures = util.KnownFolderPath(
		windows.FOLDERID_Pictures,
		nil,
		[]string{filepath.Join(home, "Pictures")},
	)
	kf.videos = util.KnownFolderPath(
		windows.FOLDERID_Videos,
		nil,
		[]string{filepath.Join(home, "Videos")},
	)
	kf.templates = util.KnownFolderPath(
		windows.FOLDERID_Templates,
		nil,
		[]string{filepath.Join(kf.roamingAppData, "Microsoft", "Windows", "Templates")},
	)
	kf.public = util.KnownFolderPath(
		windows.FOLDERID_Public,
		[]string{"PUBLIC"},
		[]string{filepath.Join(kf.systemDrive, "Users", "Public")},
	)
	kf.fonts = util.KnownFolderPath(
		windows.FOLDERID_Fonts,
		nil,
		[]string{filepath.Join(kf.systemRoot, "Fonts")},
	)
	kf.programs = util.KnownFolderPath(
		windows.FOLDERID_Programs,
		nil,
		[]string{filepath.Join(kf.roamingAppData, "Microsoft", "Windows", "Start Menu", "Programs")},
	)
	kf.commonPrograms = util.KnownFolderPath(
		windows.FOLDERID_CommonPrograms,
		nil,
		[]string{filepath.Join(kf.programData, "Microsoft", "Windows", "Start Menu", "Programs")},
	)

	return kf
}
