package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThinLongRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.5 15.5L22 12l-3.5-3.5l-.707.707l2.293 2.293H2v1h18.086l-2.293 2.293l.707.707Z"/>`),
		g.Group(children),
	)
}