package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableInsertColumnTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 3.5a.5.5 0 0 0-1 0v13a.5.5 0 0 0 1 0v-13Zm13 0a.5.5 0 0 0-1 0v13a.5.5 0 0 0 1 0v-13ZM9 3a2 2 0 0 0-2 2v2h6V5a2 2 0 0 0-2-2H9Zm-2 9V8h6v4H7Zm0 1h6v2a2 2 0 0 1-2 2H9a2 2 0 0 1-2-2v-2Z"/>`),
		g.Group(children),
	)
}