package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileMinusOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m10.5.5l.354-.354L10.707 0H10.5v.5Zm3 3h.5v-.207l-.146-.147l-.354.354Zm-1 10.5h-10v1h10v-1ZM2 13.5v-12H1v12h1ZM2.5 1h8V0h-8v1ZM13 3.5v10h1v-10h-1ZM10.146.854l3 3l.708-.708l-3-3l-.708.708ZM2.5 14a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15v-1Zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5h1ZM5 8h5V7H5v1Z"/>`),
		g.Group(children),
	)
}