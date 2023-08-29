package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Body(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="12" cy="4" r="2" fill="currentColor"/><path fill="currentColor" d="M15 22V9h5V7H4v2h5v13h2v-7h2v7z"/>`),
		g.Group(children),
	)
}