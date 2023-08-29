package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoteSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.796 3.264a.75.75 0 1 0-1.092-1.028l-4 4.25a.75.75 0 0 0 0 1.028l4 4.25a.75.75 0 1 0 1.092-1.028L10.28 7l3.516-3.736Zm-10.5.972a.75.75 0 1 0-1.092 1.028L5.72 9l-3.516 3.736a.75.75 0 1 0 1.092 1.028l4-4.25a.75.75 0 0 0 0-1.028l-4-4.25Z"/>`),
		g.Group(children),
	)
}