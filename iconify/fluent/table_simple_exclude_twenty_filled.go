package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableSimpleExcludeTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.5 2H8v6H2V4.5A2.5 2.5 0 0 1 4.5 2ZM2 9v3.5A2.5 2.5 0 0 0 4.5 15H8V9H2Zm13-1V4.5A2.5 2.5 0 0 0 12.5 2H9v6h6Zm-4 5a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2v-4Z"/>`),
		g.Group(children),
	)
}