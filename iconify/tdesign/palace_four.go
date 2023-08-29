package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PalaceFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 1.698l4 3.334V9h4V5.032l4-3.334l4 3.334V22H2V5.032l4-3.334ZM16 9h4V5.968l-2-1.666l-2 1.666V9Zm4 2H4v9h5v-3a3 3 0 1 1 6 0v3h5v-9Zm-7 9v-3a1 1 0 1 0-2 0v3h2ZM4 9h4V5.968L6 4.302L4 5.968V9Z"/>`),
		g.Group(children),
	)
}