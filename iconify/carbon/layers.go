package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Layers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 24a.997.997 0 0 1-.474-.12l-13-7l.948-1.76L16 21.864l12.526-6.744l.948 1.76l-13 7A.997.997 0 0 1 16 24Z"/><path fill="currentColor" d="M16 30a.997.997 0 0 1-.474-.12l-13-7l.948-1.76L16 27.864l12.526-6.744l.948 1.76l-13 7A.997.997 0 0 1 16 30zm0-12a.997.997 0 0 1-.474-.12l-13-7a1 1 0 0 1 0-1.76l13-7a.998.998 0 0 1 .948 0l13 7a1 1 0 0 1 0 1.76l-13 7A.997.997 0 0 1 16 18zM5.11 10L16 15.864L26.89 10L16 4.136z"/>`),
		g.Group(children),
	)
}