package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudWordsSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 3a4 4 0 0 0-3.97 3.507A3.25 3.25 0 0 0 4.25 13h7.5a3.25 3.25 0 0 0 .22-6.493A4 4 0 0 0 8 3ZM7 6h2a.5.5 0 0 1 0 1H7a.5.5 0 0 1 0-1ZM4 9.5a.5.5 0 0 1 .5-.5h2a.5.5 0 0 1 0 1h-2a.5.5 0 0 1-.5-.5ZM9.5 9h2a.5.5 0 0 1 0 1h-2a.5.5 0 0 1 0-1Z"/>`),
		g.Group(children),
	)
}