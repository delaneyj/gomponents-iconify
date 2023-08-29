package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WrapDisabled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 7H3V5h13v2M3 19h13v-2H3v2m19-7l-4-3v2H3v2h15v2l4-3Z"/>`),
		g.Group(children),
	)
}