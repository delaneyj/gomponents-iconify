package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ControlButtonTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 6.75A2.75 2.75 0 0 1 4.75 4h10.5A2.75 2.75 0 0 1 18 6.75v6.5A2.75 2.75 0 0 1 15.25 16H4.75A2.75 2.75 0 0 1 2 13.25v-6.5ZM6 7a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h.5a.5.5 0 0 0 0-1H6a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1h.5a.5.5 0 0 0 0-1H6Zm3 .5a.5.5 0 0 0-1 0V9h-.5a.5.5 0 0 0 0 1H8v1.75c0 .69.56 1.25 1.25 1.25h.25a.5.5 0 0 0 0-1h-.25a.25.25 0 0 1-.25-.25V10h.5a.5.5 0 0 0 0-1H9V7.5Zm7 0a.5.5 0 0 0-1 0v5a.5.5 0 0 0 1 0v-5Zm-4 3a.5.5 0 0 1 .5-.5h1a.5.5 0 0 0 0-1h-1a1.5 1.5 0 0 0-1.5 1.5v2a.5.5 0 0 0 1 0v-2Z"/>`),
		g.Group(children),
	)
}