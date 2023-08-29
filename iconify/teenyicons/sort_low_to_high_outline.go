package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortLowToHighOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M11.5.5h.5a.5.5 0 0 0-.5-.5v.5Zm-7 0l.354-.354a.5.5 0 0 0-.708 0L4.5.5ZM10 1h1.5V0H10v1Zm1-.5v6h1v-6h-1ZM10 7h3V6h-3v1Zm1.578 1H11v1h.578V8Zm1.396 3.658l.393-1.176l-.95-.317l-.391 1.177l.948.316ZM11 12h1.5v-1H11v1Zm.974 2.658l1-3l-.948-.316l-1 3l.948.316ZM9 10a2 2 0 0 0 2 2v-1a1 1 0 0 1-1-1H9Zm2-2a2 2 0 0 0-2 2h1a1 1 0 0 1 1-1V8Zm.578 1a.885.885 0 0 1 .84 1.165l.949.317A1.885 1.885 0 0 0 11.578 8v1ZM4.146.146l-3 3l.708.708l3-3l-.708-.708Zm0 .708l3 3l.708-.708l-3-3l-.708.708ZM4 .5V15h1V.5H4Z"/>`),
		g.Group(children),
	)
}