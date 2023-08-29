package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoEncryptionRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 17.15L10.85 8H15V6q0-1.25-.862-2.125T12.025 3q-1.25 0-2.1.863T9.075 6v.225L7.25 4.4q.55-1.525 1.875-2.463t2.9-.937Q14.1 1 15.55 2.463T17 6v2h1q.825 0 1.413.588T20 10v7.15ZM6 22q-.825 0-1.413-.588T4 20V10q0-.65.363-1.163T5.3 8.1L1.4 4.2q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l18.4 18.4q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275l-.85-.85q-.225.125-.45.188T18 22H6Zm4.6-8.55q-.275.275-.425.638t-.15.787q0 .825.588 1.412t1.412.588q.425 0 .788-.15t.637-.425l-2.85-2.85Z"/>`),
		g.Group(children),
	)
}