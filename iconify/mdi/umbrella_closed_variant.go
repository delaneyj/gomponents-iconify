package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UmbrellaClosedVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15 7l-2-1V3.5c0-.3.2-.5.5-.5s.5.2.5.5V4h2v-.5C16 2.1 14.9 1 13.5 1S11 2.1 11 3.5V6L9 7L6 6l5.1 15.3c.3.9 1.6.9 1.9 0C14.6 16.2 18 6 18 6l-3 1m-4 8L9 8.9l.7.2l1.3-1V15m2-6.9l1.3 1l.7-.2l-2 6.1V8.1Z"/>`),
		g.Group(children),
	)
}