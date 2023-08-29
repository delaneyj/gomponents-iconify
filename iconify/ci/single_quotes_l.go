package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SingleQuotesL(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 12v3.4c0 .56 0 .84.109 1.054a1 1 0 0 0 .437.437c.214.109.494.109 1.053.109h1.803c.559 0 .838 0 1.052-.109c.188-.096.341-.25.437-.437C15 16.24 15 15.96 15 15.4v-1.803c0-.559 0-.838-.109-1.052a1 1 0 0 0-.437-.437C14.24 12 13.96 12 13.4 12H10Zm0 0v-2a3 3 0 0 1 3-3"/>`),
		g.Group(children),
	)
}