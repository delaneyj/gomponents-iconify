package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BringToFront(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 10h-6V4a2.002 2.002 0 0 0-2-2H4a2.002 2.002 0 0 0-2 2v16a2.002 2.002 0 0 0 2 2h6v6a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V12a2 2 0 0 0-2-2ZM4 20L3.998 4H20v6h-8a2 2 0 0 0-2 2v8Z"/>`),
		g.Group(children),
	)
}