package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BagChecked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 19H2v2h20v-2M4 15c0 .5.2 1 .6 1.4c.4.4.9.6 1.4.6V6c-.5 0-1 .2-1.4.6C4.2 7 4 7.5 4 8v7m9.5-9h-3c0-.4.1-.8.4-1.1c.3-.3.6-.4 1.1-.4c.4 0 .8.1 1.1.4c.2.3.4.7.4 1.1M7 6v11h10V6h-2c0-.8-.3-1.5-.9-2.1S12.8 3 12 3s-1.5.3-2.1.9C9.3 4.5 9 5.2 9 6H7m11 11c.5 0 1-.2 1.4-.6c.4-.4.6-.9.6-1.4V8c0-.5-.2-1-.6-1.4c-.4-.4-.9-.6-1.4-.6v11Z"/>`),
		g.Group(children),
	)
}