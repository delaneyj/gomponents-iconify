package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionUTurnFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" d="M26.586 17.586L22 22.172V14a8 8 0 0 0-16 0v12h2V14a6 6 0 0 1 12 0v8.172l-4.586-4.586L14 19l7 7l7-7Z"/><path fill="currentColor" d="M28 2H4a2 2 0 0 0-2 2v24a2 2 0 0 0 2 2h24a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2Zm-7 24l-7-7l1.414-1.414L20 22.172V14a6 6 0 0 0-12 0v12H6V14a8 8 0 0 1 16 0v8.172l4.586-4.586L28 19Z"/>`),
		g.Group(children),
	)
}