package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrintConnectOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 10h16H4Zm2 11v-4H2v-6q0-1.275.875-2.138T5 8h14q1.275 0 2.138.863T22 11v.75q-.45-.25-.95-.425t-1.05-.25V11q0-.425-.288-.713T19 10H5q-.425 0-.713.288T4 11v4h2v-2h8.5q-.4.425-.725.925T13.25 15H8v4h5.1q.175.55.425 1.05t.6.95H6Zm11.95-.825L15.1 17.35l1.425-1.4l1.425 1.4l3.525-3.525l1.425 1.4l-4.95 4.95ZM16 8V5H8v3H6V3h12v5h-2Z"/>`),
		g.Group(children),
	)
}