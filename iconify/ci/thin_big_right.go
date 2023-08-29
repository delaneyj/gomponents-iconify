package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThinBigRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.5 18.5L24 12l-6.5-6.5l-.707.707l5.293 5.293H0v1h22.086l-5.293 5.293l.707.707Z"/>`),
		g.Group(children),
	)
}