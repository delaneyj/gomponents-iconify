package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stairs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 5v4h-4v4H7v4H3v3h7v-4h4v-4h4V8h4V5h-7Z"/>`),
		g.Group(children),
	)
}