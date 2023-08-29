package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MuseumTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMuseumTwo0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M4 44h40"/><path fill="#555" d="M8 7.273S15 4 24 4s16 3.273 16 3.273V13H8V7.273Z"/><path stroke-linecap="round" d="M10 13v25m7-25v25m7-25v25m7-25v25m7-25v25"/><path d="M7 38h34v6H7z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMuseumTwo0)"/>`),
		g.Group(children),
	)
}