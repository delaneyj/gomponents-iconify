package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandExtendedOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 15v1l-8 2.5l-7-1.9V18H1V7h8l6.2 2.3c1 .4 1.8 1.5 1.8 2.7h2c1.7 0 3 1.3 3 3M5 16V9H3v7h2m14.9-1.4c-.2-.4-.5-.6-.9-.6h-5.4c-.5 0-1.1-.1-1.6-.2L9.7 13l.6-1.9l2.4.8c.3 0 2.3.1 2.3.1c0-.4-.2-.7-.6-.8L8.6 9H7v5.5l7 1.9l5.9-1.8Z"/>`),
		g.Group(children),
	)
}