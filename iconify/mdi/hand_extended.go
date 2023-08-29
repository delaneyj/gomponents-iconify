package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandExtended(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 13c.6 0 1.1.2 1.5.6c.3.4.5.9.5 1.4l-8 3l-7-2V7h1.9l7.3 2.7c.5.2.8.6.8 1.1c0 .3-.1.6-.3.8s-.6.4-.9.4H13l-1.8-.7l-.3.9l2.1.8h7M1 7h4v11H1V7Z"/>`),
		g.Group(children),
	)
}