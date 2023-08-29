package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gamepad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.5 9l-3 3l3 3H22V9M9 16.5V22h6v-5.5l-3-3M7.5 9H2v6h5.5l3-3M15 7.5V2H9v5.5l3 3l3-3Z"/>`),
		g.Group(children),
	)
}