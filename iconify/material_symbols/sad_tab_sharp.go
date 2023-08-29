package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SadTabSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 16.5H9V15h6v1.5h1.5V15H15v-1.5H9V15H7.5ZM8 12h1.5V9.5H8Zm6.5 0H16V9.5h-1.5ZM2 20V4h8l2 2h10v14Z"/>`),
		g.Group(children),
	)
}