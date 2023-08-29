package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M20 4h-8v6h8V4Zm0 8h-8v8h8v-8Zm2 8v-8h6v8h-6Zm-2 2h-8v6h8v-6Zm2 6v-6h6v1a5 5 0 0 1-5 5h-1Zm0-18V4h1a5 5 0 0 1 5 5v1h-6ZM9 4h1v6H4V9a5 5 0 0 1 5-5Zm-5 8h6v8H4v-8Zm0 10h6v6H9a5 5 0 0 1-5-5v-1Z"/>`),
		g.Group(children),
	)
}