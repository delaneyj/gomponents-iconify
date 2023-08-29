package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoEncryption(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 17.15L10.85 8H15V6q0-1.25-.862-2.125T12.025 3q-1.25 0-2.1.875T9.075 6v.225L7.25 4.4Q7.8 2.875 9.1 1.937T12.025 1Q14.1 1 15.55 2.463T17 6v2h1q.825 0 1.413.588T20 10v7.15Zm.5 6.15l-1.55-1.55q-.225.125-.463.188T18 22H6q-.825 0-1.413-.588T4 20V10q0-.65.363-1.163T5.3 8.1L.7 3.5l1.4-1.4l19.8 19.8l-1.4 1.4Zm-9.875-9.875q-.275.3-.425.663t-.15.762q0 .825.587 1.413t1.413.587q.4 0 .763-.15t.662-.425l-2.85-2.85Z"/>`),
		g.Group(children),
	)
}