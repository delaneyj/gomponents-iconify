package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextColumnTwoLeftTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21 6a1 1 0 0 0-1-1h-9a1 1 0 1 0 0 2h9a1 1 0 0 0 1-1ZM8 6a1 1 0 0 0-1-1H4a1 1 0 0 0 0 2h3a1 1 0 0 0 1-1Zm13 4a1 1 0 0 0-1-1h-9a1 1 0 1 0 0 2h9a1 1 0 0 0 1-1ZM8 10a1 1 0 0 0-1-1H4a1 1 0 0 0 0 2h3a1 1 0 0 0 1-1Zm13 4a1 1 0 0 0-1-1h-9a1 1 0 1 0 0 2h9a1 1 0 0 0 1-1ZM8 14a1 1 0 0 0-1-1H4a1 1 0 1 0 0 2h3a1 1 0 0 0 1-1Zm13 4a1 1 0 0 0-1-1h-9a1 1 0 1 0 0 2h9a1 1 0 0 0 1-1ZM8 18a1 1 0 0 0-1-1H4a1 1 0 1 0 0 2h3a1 1 0 0 0 1-1Z"/>`),
		g.Group(children),
	)
}