package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddHome(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.5 19h1v-2.5H21v-1h-2.5V13h-1v2.5H15v1h2.5V19Zm.5 2q-2.075 0-3.538-1.463T13 16q0-2.075 1.463-3.538T18 11q2.075 0 3.538 1.463T23 16q0 2.075-1.463 3.538T18 21ZM4 19V7l8-6l8 6v2.3q-.5-.15-1-.225T18 9q-2.925 0-4.963 2.038T11 16q0 .8.163 1.55t.512 1.45H4Z"/>`),
		g.Group(children),
	)
}