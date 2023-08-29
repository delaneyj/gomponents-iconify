package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HdrOnSelectOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.5 22v-2h-2v-1.5h2v-2H22v2h2V20h-2v2h-1.5ZM12 12q1.65 0 2.825-1.175T16 8q0-1.65-1.175-2.825T12 4q-1.65 0-2.825 1.175T8 8q0 1.65 1.175 2.825T12 12Zm0 2q-2.5 0-4.25-1.75T6 8q0-2.5 1.75-4.25T12 2q2.5 0 4.25 1.75T18 8q0 2.5-1.75 4.25T12 14Zm1 8v-6h5v3.9h-.9L18 22h-1.5l-.9-2h-1.1v2H13Zm1.5-3.5h2v-1h-2v1ZM0 22v-6h1.5v2h2v-2H5v6H3.5v-2.5h-2V22H0Zm6.5 0v-6H10q.6 0 1.05.45t.45 1.05v3q0 .6-.45 1.05T10 22H6.5ZM8 20.5h2v-3H8v3Z"/>`),
		g.Group(children),
	)
}