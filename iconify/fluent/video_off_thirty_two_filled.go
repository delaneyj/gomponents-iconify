package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoOffThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m20.5 21.914l7.793 7.793a1 1 0 0 0 1.414-1.414l-26-26a1 1 0 0 0-1.414 1.414L4.536 5.95A4.5 4.5 0 0 0 2 10v12a4.5 4.5 0 0 0 4.5 4.5H16a4.5 4.5 0 0 0 4.5-4.5v-.086Zm1.5-3.449l5.245 5.244c1.207.664 2.755-.198 2.755-1.642V9.933c0-1.538-1.756-2.416-2.987-1.493L22 12.2v6.265ZM9.035 5.5L20.5 16.965V10A4.5 4.5 0 0 0 16 5.5H9.035Z"/>`),
		g.Group(children),
	)
}