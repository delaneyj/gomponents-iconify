package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortHighToLowOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M11.5.5h.5a.5.5 0 0 0-.5-.5v.5Zm1 11l.474.158l-.474-.158Zm.482-1.446l-.474-.158l.474.158ZM4.5 14.5l-.354.354a.5.5 0 0 0 .708 0L4.5 14.5ZM10 1h1.5V0H10v1Zm1-.5v6h1v-6h-1ZM10 7h3V6h-3v1Zm1.862 1H11v1h.862V8Zm1.112 3.658l.482-1.446l-.948-.316l-.482 1.446l.948.316ZM11 12h1.5v-1H11v1Zm.974 2.658l1-3l-.948-.316l-1 3l.948.316ZM9 10a2 2 0 0 0 2 2v-1a1 1 0 0 1-1-1H9Zm2-2a2 2 0 0 0-2 2h1a1 1 0 0 1 1-1V8Zm.862 1a.68.68 0 0 1 .646.896l.948.316A1.68 1.68 0 0 0 11.862 8v1Zm-7.008 5.854l3-3l-.708-.708l-3 3l.708.708Zm0-.708l-3-3l-.708.708l3 3l.708-.708ZM4 0v14.5h1V0H4Z"/>`),
		g.Group(children),
	)
}