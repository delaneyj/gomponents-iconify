package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FridgeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.998 12v10a1 1 0 0 1-1 1h-14a1 1 0 0 1-1-1V12h16Zm-11 2h-2v5h2v-5Zm10-13a1 1 0 0 1 1 1v8h-16V2a1 1 0 0 1 1-1h14Zm-10 3h-2v4h2V4Z"/>`),
		g.Group(children),
	)
}