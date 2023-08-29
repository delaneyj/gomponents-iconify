package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SelectObjectTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7 5a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm1 0a1 1 0 0 0 1 1h6a1 1 0 1 0 0-2H9a1 1 0 0 0-1 1ZM5 8a1 1 0 0 0-1 1v6a1 1 0 1 0 2 0V9a1 1 0 0 0-1-1Zm14 0a1 1 0 0 0-1 1v6a1 1 0 1 0 2 0V9a1 1 0 0 0-1-1ZM9 20a1 1 0 1 1 0-2h6a1 1 0 1 1 0 2H9Zm-4 1a2 2 0 1 0 0-4a2 2 0 0 0 0 4ZM21 5a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm-2 16a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/>`),
		g.Group(children),
	)
}