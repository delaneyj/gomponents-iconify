package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VectorPointEdit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 9v6h6V9H9m2 2h2v2h-2v-2m10.2 2c-.1 0-.3.1-.4.2l-1 1l2.1 2.1l1-1c.2-.2.2-.6 0-.8l-1.3-1.3c-.2-.1-.3-.2-.4-.2m-2.1 1.8L13 20.9V23h2.1l6.1-6.2l-2.1-2Z"/>`),
		g.Group(children),
	)
}