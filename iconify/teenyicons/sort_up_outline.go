package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortUpOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m3.5.5l.354-.354a.5.5 0 0 0-.708 0L3.5.5ZM3.146.146l-3 3l.708.708l3-3l-.708-.708Zm0 .708l3 3l.708-.708l-3-3l-.708.708ZM3 .5V15h1V.5H3ZM9 4h6V3H9v1Zm0 4h4V7H9v1Zm0 4h2v-1H9v1Z"/>`),
		g.Group(children),
	)
}