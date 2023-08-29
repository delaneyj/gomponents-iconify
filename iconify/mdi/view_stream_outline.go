package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewStreamOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 6v12h17V6H4m15 10H6v-3h13v3M6 11V8h13v3H6Z"/>`),
		g.Group(children),
	)
}