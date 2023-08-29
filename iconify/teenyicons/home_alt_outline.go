package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeAltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m7.5.5l.354-.354a.5.5 0 0 0-.708 0L7.5.5Zm-6 6l-.354-.354L1 6.293V6.5h.5Zm12 0h.5v-.207l-.146-.147l-.354.354Zm.354-.354l-6-6l-.708.708l6 6l.708-.708Zm-6.708-6l-6 6l.708.708l6-6l-.708-.708ZM14 13.5v-7h-1v7h1Zm-13-7v7h1v-7H1ZM2.5 15h10v-1h-10v1ZM13 13.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5h-1Zm-12 0A1.5 1.5 0 0 0 2.5 15v-1a.5.5 0 0 1-.5-.5H1Z"/>`),
		g.Group(children),
	)
}