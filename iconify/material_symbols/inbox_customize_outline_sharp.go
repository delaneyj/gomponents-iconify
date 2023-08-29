package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InboxCustomizeOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 18h6.325q.175.55.4 1.05t.55.95H3V2h18v7.65q-.475-.225-.975-.363T19 9.075V4H5v9h4.2q.225.675.75 1.175t1.175.7q-.075.5-.1 1.012t.05 1.013q-.9-.175-1.687-.663T8 15H5v3Zm0 0h6.325H5Zm12.025 3l-.3-1.5q-.3-.125-.563-.263t-.537-.337l-1.45.45l-1-1.7l1.15-1q-.05-.3-.05-.65t.05-.65l-1.15-1l1-1.7l1.45.45q.275-.2.537-.337t.563-.263l.3-1.5h2l.3 1.5q.3.125.563.263t.537.337l1.45-.45l1 1.7l-1.15 1q.05.3.05.65t-.05.65l1.15 1l-1 1.7l-1.45-.45q-.275.2-.537.338t-.563.262l-.3 1.5h-2Zm1-3q.825 0 1.413-.588T20.024 16q0-.825-.587-1.413T18.025 14q-.825 0-1.412.588T16.024 16q0 .825.588 1.413t1.412.587Z"/>`),
		g.Group(children),
	)
}