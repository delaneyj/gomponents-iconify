package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlertTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 1a4 4 0 0 0-4 4v1.382l-.947 1.894A.5.5 0 0 0 1.5 9h9a.5.5 0 0 0 .447-.724L10 6.382V5a4 4 0 0 0-4-4Zm0 10.5A2 2 0 0 1 4.063 10h3.874A2 2 0 0 1 6 11.5Z"/>`),
		g.Group(children),
	)
}