package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Diagram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M13 30a11 11 0 0 1 0-22a1 1 0 0 1 1 1v9h9a1 1 0 0 1 1 1a11 11 0 0 1-11 11Zm-1-19.94A9 9 0 1 0 21.94 20H14a2 2 0 0 1-2-2Z"/><path fill="currentColor" d="M28 14h-9a2 2 0 0 1-2-2V3a1 1 0 0 1 1-1a11 11 0 0 1 11 11a1 1 0 0 1-1 1Zm-9-2h7.94A9 9 0 0 0 19 4.06Z"/>`),
		g.Group(children),
	)
}