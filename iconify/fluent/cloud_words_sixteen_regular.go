package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudWordsSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 7a3 3 0 0 1 6 0a.5.5 0 0 0 .5.5h.25a2.25 2.25 0 0 1 0 4.5h-7.5a2.25 2.25 0 0 1 0-4.5h.25A.5.5 0 0 0 5 7Zm3-4a4 4 0 0 0-3.97 3.507A3.25 3.25 0 0 0 4.25 13h7.5a3.25 3.25 0 0 0 .22-6.493A4 4 0 0 0 8 3ZM7 6a.5.5 0 0 0 0 1h2a.5.5 0 0 0 0-1H7ZM4 9.5a.5.5 0 0 1 .5-.5h2a.5.5 0 0 1 0 1h-2a.5.5 0 0 1-.5-.5ZM9.5 9a.5.5 0 0 0 0 1h2a.5.5 0 0 0 0-1h-2Z"/>`),
		g.Group(children),
	)
}