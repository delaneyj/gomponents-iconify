package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageBreakOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M.5 8.993H0v1h.5v-1Zm2 1H3v-1h-.5v1Zm2-1H4v1h.5v-1Zm2 1H7v-1h-.5v1Zm2-1H8v1h.5v-1Zm2 1h.5v-1h-.5v1Zm2-1H12v1h.5v-1Zm2 1h.5v-1h-.5v1ZM10.5.5l.354-.354L10.707 0H10.5v.5Zm3 3h.5v-.207l-.146-.147l-.354.354ZM.5 9.993h2v-1h-2v1Zm4 0h2v-1h-2v1Zm4 0h2v-1h-2v1Zm4 0h2v-1h-2v1Zm0 4.007h-10v1h10v-1ZM2.5 1h8V0h-8v1Zm7.646-.146l3 3l.708-.708l-3-3l-.708.708ZM13 3.5V8h1V3.5h-1Zm0 8.25v1.75h1v-1.75h-1ZM2 8V1.5H1V8h1Zm0 5.5V11H1v2.5h1Zm.5.5a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15v-1Zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM2.5 0A1.5 1.5 0 0 0 1 1.5h1a.5.5 0 0 1 .5-.5V0Z"/>`),
		g.Group(children),
	)
}