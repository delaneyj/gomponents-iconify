package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareScreenTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M19.75 4A2.25 2.25 0 0 1 22 6.25v11.5A2.25 2.25 0 0 1 19.75 20H4.25A2.25 2.25 0 0 1 2 17.75V6.25A2.25 2.25 0 0 1 4.25 4zM12 7.245a.75.75 0 0 0-.53.22L8.22 10.72a.75.75 0 0 0 1.06 1.06l1.97-1.972v6.445a.75.75 0 1 0 1.5 0V9.806l1.974 1.974a.75.75 0 1 0 1.06-1.06L12.53 7.465a.75.75 0 0 0-.53-.22z" fill="currentColor" fill-rule="nonzero"/>`),
		g.Group(children),
	)
}