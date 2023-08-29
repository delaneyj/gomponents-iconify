package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoBackpackOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20 17.175l-2-2V8q0-.825-.588-1.413T16 6H8.825L7 4.15V2h3v2h4V2h3v2.15q1.3.35 2.15 1.4T20 8v9.175Zm-3.5-3.5L14.825 12H16.5v1.675ZM11.175 12v2H7.5v-2h3.675Zm2.375-1.275Zm-2 3.65Zm-6.375-9.2l1.4 1.4q-.275.275-.425.638T6 8v12h12v-1.975L20 20v2H4V8q0-.825.313-1.55t.862-1.275Zm14.6 17.425L1.4 4.225L2.8 2.8l18.375 18.4l-1.4 1.4Z"/>`),
		g.Group(children),
	)
}