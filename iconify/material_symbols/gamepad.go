package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gamepad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 10.5l-3-3V2h6v5.5l-3 3Zm4.5 4.5l-3-3l3-3H22v6h-5.5ZM2 15V9h5.5l3 3l-3 3H2Zm7 7v-5.5l3-3l3 3V22H9Z"/>`),
		g.Group(children),
	)
}