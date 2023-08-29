package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 5v8h10.17l-2.58-2.59L15 9l5 5l-5 5l-1.41-1.41L16.17 15H4V5h2Z"/>`),
		g.Group(children),
	)
}