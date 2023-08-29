package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareScreenTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M19.75 4A2.25 2.25 0 0 1 22 6.25v11.5A2.25 2.25 0 0 1 19.75 20H4.25A2.25 2.25 0 0 1 2 17.75V6.25A2.25 2.25 0 0 1 4.25 4h15.5zm0 1.5H4.25a.75.75 0 0 0-.75.75v11.5c0 .414.336.75.75.75h15.5a.75.75 0 0 0 .75-.75V6.25a.75.75 0 0 0-.75-.75zM12 7.245a.75.75 0 0 1 .53.22l3.255 3.255a.75.75 0 1 1-1.06 1.06L12.75 9.806v6.447a.75.75 0 0 1-1.5 0V9.808L9.28 11.78a.75.75 0 1 1-1.06-1.06l3.25-3.254a.75.75 0 0 1 .53-.22z" fill="currentColor" fill-rule="nonzero"/>`),
		g.Group(children),
	)
}