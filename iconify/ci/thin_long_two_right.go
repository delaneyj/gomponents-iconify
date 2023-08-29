package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThinLongTwoRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16 15l3-3l-3-3l-.707.707l1.793 1.793H5v1h12.086l-1.793 1.793L16 15Z"/>`),
		g.Group(children),
	)
}