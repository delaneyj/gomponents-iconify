package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RibbonB(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M256 72c-48.523 0-88 39.477-88 88s39.477 88 88 88 88-39.477 88-88-39.477-88-88-88z" fill="currentColor"/><path d="M256 32c-70.692 0-128 57.308-128 128s57.308 128 128 128 128-57.308 128-128S326.692 32 256 32zm0 224c-53.02 0-96-42.98-96-96s42.98-96 96-96 96 42.98 96 96-42.98 96-96 96z" fill="currentColor"/><g><path d="M147.092 254.21L64 400h96l48 80 48-105.807 33.641-74.154A144.328 144.328 0 0 1 256 304c-43.505 0-82.503-19.293-108.908-49.79z" fill="currentColor"/><path d="M364.908 254.211c-15.077 17.412-34.26 31.172-56.043 39.774l-44.752 98.092L304 480l48-80h96l-83.092-145.789z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}