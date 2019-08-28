package log

type WriterType int

const (
	_stdout WriterType = iota
	_file
)

var Conf = &Config{}

var writerTypes = map[string]WriterType{
	"stdout": _stdout,
	"file":   _file,
}

type WriterConfig struct {
	Type   string
	Path   string
	Level  string
	enable bool
}

type Config struct {
	Writers     []WriterConfig
	ServiceName string
}
