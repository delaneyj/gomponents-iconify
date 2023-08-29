package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckerboardMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 17v2h8v-2M8 16h4v-4H8v4m4-4h4V8h-4v4M2 2v20h11.5c-.5-.6-.9-1.3-1.2-2H8v-4H4v-4h4V8H4V4h4v4h4V4h4v4h4v4.4c.7.3 1.4.7 2 1.2V2H2Z"/>`),
		g.Group(children),
	)
}