package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Exclude(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.54 7.54h-1a1 1 0 0 0-1 1v1a1 1 0 0 0 2 0a1 1 0 1 0 0-2Zm5.92 5.92a1 1 0 0 0-1 1a1 1 0 0 0 0 2h1a1 1 0 0 0 1-1v-1a1 1 0 0 0-1-1ZM21 7.54h-4.54V3a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v12.46a1 1 0 0 0 1 1h4.54V21a1 1 0 0 0 1 1H21a1 1 0 0 0 1-1V8.54a1 1 0 0 0-1-1ZM20 20H9.54v-3.54a1 1 0 0 0 0-2a1 1 0 0 0-2 0H4V4h10.46v3.54a1 1 0 0 0 0 2a1 1 0 0 0 2 0H20Z"/>`),
		g.Group(children),
	)
}