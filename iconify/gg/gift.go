package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gift(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17.535 2.879a3 3 0 0 0-4.242 0l-1.414 1.414c-.06.06-.118.122-.172.186a3.01 3.01 0 0 0-.171-.186L10.12 2.879A3 3 0 1 0 5.88 7.12l.877.88H1v6h2v8h18v-8h2V8h-6.343l.878-.879a3 3 0 0 0 0-4.242ZM14.707 7.12l1.414-1.414a1 1 0 0 0-1.414-1.414l-1.414 1.414a1 1 0 0 0 1.414 1.414Zm-4.586-1.414L8.707 4.293a1 1 0 1 0-1.414 1.414l1.414 1.414a1 1 0 1 0 1.414-1.414ZM21 10v2H3v-2h18Zm-8.083 4H19v6h-6.083v-6Zm-1.834 0v6H5v-6h6.083Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}