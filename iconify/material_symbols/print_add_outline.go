package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrintAddOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 21v-3h-3v-2h3v-3h2v3h3v2h-3v3h-2ZM4 10h16H4Zm2 11v-4H2v-6q0-1.275.875-2.138T5 8h14q1.275 0 2.138.863T22 11v.8q-.45-.25-.95-.438T20 11.1q0-.425-.288-.762T19 10H5q-.425 0-.713.288T4 11v4h2v-2h8.55q-.4.425-.7.925T13.35 15H8v4h5.35q.175.55.5 1.05t.7.95H6ZM16 8V5H8v3H6V3h12v5h-2Z"/>`),
		g.Group(children),
	)
}