package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableBottomRowThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 4a5 5 0 0 0-5 5v11h2V9a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v11h2V9a5 5 0 0 0-5-5H9Zm11 18v6h-8v-6h8Zm3 6h-1v-6h6v1a5 5 0 0 1-5 5Zm-13-6v6H9a5 5 0 0 1-5-5v-1h6Z"/>`),
		g.Group(children),
	)
}