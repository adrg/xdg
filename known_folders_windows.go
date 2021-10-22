package xdg

import (
	"os"
	"path/filepath"
	"strings"

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
	kf.systemDrive = strings.TrimRight(knownFolder(
		nil,
		[]string{"SystemDrive"},
		[]string{filepath.VolumeName(home), `C:\`},
	), sep) + sep
	kf.systemRoot = knownFolder(
		windows.FOLDERID_Windows,
		[]string{"SystemRoot", "windir"},
		[]string{filepath.Join(kf.systemDrive, "Windows")},
	)
	kf.programData = knownFolder(
		windows.FOLDERID_ProgramData,
		[]string{"ALLUSERSPROFILE", "PROGRAMDATA"},
		[]string{filepath.Join(kf.systemDrive, "ProgramData")},
	)
	kf.roamingAppData = knownFolder(
		windows.FOLDERID_RoamingAppData,
		[]string{"APPDATA"},
		[]string{filepath.Join(home, "AppData", "Roaming")},
	)
	kf.localAppData = knownFolder(
		windows.FOLDERID_LocalAppData,
		[]string{"LOCALAPPDATA"},
		[]string{filepath.Join(home, "AppData", "Local")},
	)
	kf.desktop = knownFolder(
		windows.FOLDERID_Desktop,
		nil,
		[]string{filepath.Join(home, "Desktop")},
	)
	kf.downloads = knownFolder(
		windows.FOLDERID_Downloads,
		nil,
		[]string{filepath.Join(home, "Downloads")},
	)
	kf.documents = knownFolder(
		windows.FOLDERID_Documents,
		nil,
		[]string{filepath.Join(home, "Documents")},
	)
	kf.music = knownFolder(
		windows.FOLDERID_Music,
		nil,
		[]string{filepath.Join(home, "Music")},
	)
	kf.pictures = knownFolder(
		windows.FOLDERID_Pictures,
		nil,
		[]string{filepath.Join(home, "Pictures")},
	)
	kf.videos = knownFolder(
		windows.FOLDERID_Videos,
		nil,
		[]string{filepath.Join(home, "Videos")},
	)
	kf.templates = knownFolder(
		windows.FOLDERID_Templates,
		nil,
		[]string{filepath.Join(kf.roamingAppData, "Microsoft", "Windows", "Templates")},
	)
	kf.public = knownFolder(
		windows.FOLDERID_Public,
		[]string{"PUBLIC"},
		[]string{filepath.Join(kf.systemDrive, "Users", "Public")},
	)
	kf.fonts = knownFolder(
		windows.FOLDERID_Fonts,
		nil,
		[]string{filepath.Join(kf.systemRoot, "Fonts")},
	)
	kf.programs = knownFolder(
		windows.FOLDERID_Programs,
		nil,
		[]string{filepath.Join(kf.roamingAppData, "Microsoft", "Windows", "Start Menu", "Programs")},
	)
	kf.commonPrograms = knownFolder(
		windows.FOLDERID_CommonPrograms,
		nil,
		[]string{filepath.Join(kf.programData, "Microsoft", "Windows", "Start Menu", "Programs")},
	)

	return kf
}

func knownFolder(id *windows.KNOWNFOLDERID, envVars []string, fallbacks []string) string {
	if id != nil {
		flags := []uint32{windows.KF_FLAG_DEFAULT, windows.KF_FLAG_DEFAULT_PATH}
		for _, flag := range flags {
			if p, _ := windows.KnownFolderPath(id, flag|windows.KF_FLAG_DONT_VERIFY); p != "" {
				return p
			}
		}
	}

	for _, envVar := range envVars {
		if p := os.Getenv(envVar); p != "" {
			return p
		}
	}

	for _, fallback := range fallbacks {
		if fallback != "" {
			return fallback
		}
	}

	return ""
}
