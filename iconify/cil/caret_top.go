package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M496 400.357H16.333v-36.449L256.047 96L496 365.81Zm-440.708-32h400.149L255.975 144.07Z"/>`),
		g.Group(children),
	)
}