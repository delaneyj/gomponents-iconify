package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareScreenTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M23.75 5A2.25 2.25 0 0 1 26 7.25v13.5A2.25 2.25 0 0 1 23.75 23H4.25A2.25 2.25 0 0 1 2 20.75V7.25A2.25 2.25 0 0 1 4.25 5h19.5zm0 1.5H4.25a.75.75 0 0 0-.75.75v13.5c0 .414.336.75.75.75h19.5a.75.75 0 0 0 .75-.75V7.25a.75.75 0 0 0-.75-.75zM13.998 8.62a.75.75 0 0 1 .53.22l3.255 3.254a.75.75 0 0 1-1.061 1.06l-1.974-1.973v7.446a.75.75 0 1 1-1.5 0v-7.444l-1.97 1.972a.75.75 0 1 1-1.061-1.06l3.25-3.255a.75.75 0 0 1 .53-.22z" fill="currentColor" fill-rule="nonzero"/>`),
		g.Group(children),
	)
}