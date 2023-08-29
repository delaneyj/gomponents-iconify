package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotEqual(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 10H9V8h12v2m0 6H9v-2h12v2M4 5h2v11H4V5m2 13v2H4v-2h2Z"/>`),
		g.Group(children),
	)
}