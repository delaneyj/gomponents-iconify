package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabletThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 8.25A3.25 3.25 0 0 1 5.25 5h21.5A3.25 3.25 0 0 1 30 8.25v15.5A3.25 3.25 0 0 1 26.75 27H5.25A3.25 3.25 0 0 1 2 23.75V8.25ZM13 21a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2h-6Z"/>`),
		g.Group(children),
	)
}