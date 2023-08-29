package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Widgets(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3h8v4.34l5.66-5.65l5.65 5.65L16.66 13H21v8h-8v-8h3.66L11 7.34V11H3V3m0 10h8v8H3v-8Z"/>`),
		g.Group(children),
	)
}