package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrintLockOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 21v-5h1v-1q0-.825.588-1.413T19 13q.825 0 1.413.588T21 15v1h1v5h-6Zm2-5h2v-1q0-.425-.288-.713T19 14q-.425 0-.713.288T18 15v1ZM4 10h16H4Zm2 11v-4H2V8h20v3.75q-.45-.25-.95-.425t-1.05-.25V10H4v5h2v-2h8.5q-.4.425-.725.925T13.25 15H8v4h5.1q.175.55.425 1.05t.6.95H6ZM16 8V5H8v3H6V3h12v5h-2Z"/>`),
		g.Group(children),
	)
}