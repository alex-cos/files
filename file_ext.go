package files

// nolint: varnamelen
const (
	// Documents.
	DOC  string = "doc"
	DOCX string = "docx"
	ODT  string = "odt"
	PDF  string = "pdf"
	TXT  string = "txt"
	RTF  string = "rtf"
	MD   string = "md"

	// Tableurs.
	XLS  string = "xls"
	XLSX string = "xlsx"
	ODS  string = "ods"
	CSV  string = "csv"

	// Présentations.
	PPT  string = "ppt"
	PPTX string = "pptx"
	ODP  string = "odp"

	// Images.
	JPG  string = "jpg"
	JPEG string = "jpeg"
	PNG  string = "png"
	GIF  string = "gif"
	BMP  string = "bmp"
	TIFF string = "tiff"
	SVG  string = "svg"
	WEBP string = "webp"

	// Audio.
	MP3  string = "mp3"
	WAV  string = "wav"
	FLAC string = "flac"
	AAC  string = "aac"
	OGG  string = "ogg"

	// Vidéo.
	MP4 string = "mp4"
	AVI string = "avi"
	MKV string = "mkv"
	MOV string = "mov"
	WMV string = "wmv"
	FLV string = "flv"

	// Archives.
	ZIP    string = "zip"
	RAR    string = "rar"
	TAR    string = "tar"
	GZ     string = "gz"
	SEVENZ string = "7z"

	// Exécutables.
	EXE string = "exe"
	MSI string = "msi"
	SH  string = "sh"
	BAT string = "bat"
	APP string = "app"

	// Code.
	HTML string = "html"
	CSS  string = "css"
	JS   string = "js"
	TS   string = "ts"
	PY   string = "py"
	GO   string = "go"
	CPP  string = "cpp"
	C    string = "c"
	H    string = "h"
	HH   string = "hh"
	JAVA string = "java"
	PHP  string = "php"
	SQL  string = "sql"
)

// nolint: gochecknoglobals
var FileDescriptions = map[string]string{
	// Documents.
	DOC:  "Microsoft Word Document",
	DOCX: "Microsoft Word Document (XML)",
	ODT:  "OpenDocument Text Document",
	PDF:  "Portable Document Format",
	TXT:  "Text File",
	RTF:  "Rich Text Format",
	MD:   "Markdown Document",

	// Tableurs.
	XLS:  "Microsoft Excel Spreadsheet",
	XLSX: "Microsoft Excel Spreadsheet (XML)",
	ODS:  "OpenDocument Spreadsheet",
	CSV:  "Comma-Separated Values",

	// Présentations.
	PPT:  "Microsoft PowerPoint Presentation",
	PPTX: "Microsoft PowerPoint Presentation (XML)",
	ODP:  "OpenDocument Presentation",

	// Images.
	JPG:  "JPEG Image",
	JPEG: "JPEG Image",
	PNG:  "Portable Network Graphics",
	GIF:  "Graphics Interchange Format",
	BMP:  "Bitmap Image",
	TIFF: "Tagged Image File Format",
	SVG:  "Scalable Vector Graphics",
	WEBP: "WebP Image Format",

	// Audio.
	MP3:  "MP3 Audio File",
	WAV:  "Waveform Audio File",
	FLAC: "Free Lossless Audio Codec",
	AAC:  "Advanced Audio Codec",
	OGG:  "Ogg Vorbis Audio File",

	// Vidéo.
	MP4: "MPEG-4 Video File",
	AVI: "Audio Video Interleave",
	MKV: "Matroska Video File",
	MOV: "QuickTime Movie File",
	WMV: "Windows Media Video",
	FLV: "Flash Video File",

	// Archives.
	ZIP:    "ZIP Archive",
	RAR:    "RAR Archive",
	TAR:    "TAR Archive",
	GZ:     "Gzip Compressed Archive",
	SEVENZ: "7-Zip Archive",

	// Exécutables.
	EXE: "Windows Executable File",
	MSI: "Windows Installer Package",
	SH:  "Shell Script",
	BAT: "Batch File",
	APP: "macOS Application",

	// Code.
	HTML: "HTML Document",
	CSS:  "Cascading Style Sheets",
	JS:   "JavaScript File",
	TS:   "TypeScript File",
	PY:   "Python Script",
	GO:   "Go Source Code",
	CPP:  "C++ Source Code",
	C:    "C Source Code",
	H:    "C Header",
	HH:   "C++ Header",
	JAVA: "Java Source Code",
	PHP:  "PHP Script",
	SQL:  "SQL Database Query",
}

func GetAllExtensions() []string {
	extensions := make([]string, 0, len(FileDescriptions))

	for ext := range FileDescriptions {
		extensions = append(extensions, ext)
	}

	return extensions
}

func GetFileDescription(ext string) string {
	return FileDescriptions[ext]
}
