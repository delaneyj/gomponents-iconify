package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionUTurn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26.586 19.586L22 24.172V12a8 8 0 0 0-16 0v16h2V12a6 6 0 0 1 12 0v12.172l-4.586-4.586L14 21l7 7l7-7Z"/>`),
		g.Group(children),
	)
}