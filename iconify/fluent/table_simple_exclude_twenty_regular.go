package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableSimpleExcludeTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.5 2A2.5 2.5 0 0 0 2 4.5v8A2.5 2.5 0 0 0 4.5 15h4a.5.5 0 0 0 .5-.5V9h5.5a.5.5 0 0 0 .5-.5v-4A2.5 2.5 0 0 0 12.5 2h-8ZM8 9v5H4.5A1.5 1.5 0 0 1 3 12.5V9h5Zm0-1H3V4.5A1.5 1.5 0 0 1 4.5 3H8v5Zm1 0V3h3.5A1.5 1.5 0 0 1 14 4.5V8H9Zm2 5a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2v-4Z"/>`),
		g.Group(children),
	)
}