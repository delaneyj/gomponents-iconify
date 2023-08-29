package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 16.25A7.25 7.25 0 0 1 9.25 9h3.5a1.25 1.25 0 1 1 0 2.5h-3.5a4.75 4.75 0 1 0 0 9.5h3.5a1.25 1.25 0 1 1 0 2.5h-3.5A7.25 7.25 0 0 1 2 16.25Zm28 0A7.25 7.25 0 0 0 22.75 9h-3.5a1.25 1.25 0 1 0 0 2.5h3.5a4.75 4.75 0 1 1 0 9.5h-3.5a1.25 1.25 0 1 0 0 2.5h3.5A7.25 7.25 0 0 0 30 16.25ZM9.75 15a1.25 1.25 0 1 0 0 2.5h12.5a1.25 1.25 0 1 0 0-2.5H9.75Z"/>`),
		g.Group(children),
	)
}