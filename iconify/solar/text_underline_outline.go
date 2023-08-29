package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextUnderlineOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.75 3a.75.75 0 0 0-1.5 0v6a8.75 8.75 0 1 0 17.5 0V3a.75.75 0 0 0-1.5 0v6a7.25 7.25 0 1 1-14.5 0V3ZM4 20.25a.75.75 0 0 0 0 1.5h16a.75.75 0 0 0 0-1.5H4Z"/>`),
		g.Group(children),
	)
}