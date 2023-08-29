package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditCalendarOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22V4h3V2h2v2h8V2h2v2h3v8h-2v-2H5v10h7v2H3Zm19.125-5L20 14.875l1.425-1.425l2.125 2.125L22.125 17ZM14 23v-2.125l5.3-5.3l2.125 2.125l-5.3 5.3H14ZM5 8h14V6H5v2Zm0 0V6v2Z"/>`),
		g.Group(children),
	)
}