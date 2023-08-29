package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableBottomRowThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 4a5 5 0 0 0-5 5v14a5 5 0 0 0 5 5h14a5 5 0 0 0 5-5V9a5 5 0 0 0-5-5H9ZM6 9a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v10H6V9Zm0 14v-2h5v5H9a3 3 0 0 1-3-3Zm7 3v-5h6v5h-6Zm8 0v-5h5v2a3 3 0 0 1-3 3h-2Z"/>`),
		g.Group(children),
	)
}