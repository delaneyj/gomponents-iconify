package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShortDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11 16.17l-3.59-3.58L6 14l6 6l6-6l-1.41-1.41L13 16.17V4h-2v12.17Z"/>`),
		g.Group(children),
	)
}