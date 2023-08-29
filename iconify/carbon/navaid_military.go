package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavaidMilitary(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 30a14 14 0 1 1 14-14a14.016 14.016 0 0 1-14 14Zm0-26a12 12 0 1 0 12 12A12.014 12.014 0 0 0 16 4Z"/><path fill="currentColor" d="M16 24a8 8 0 1 1 8-8a8.01 8.01 0 0 1-8 8Zm0-14a6 6 0 1 0 6 6a6.007 6.007 0 0 0-6-6Z"/><circle cx="16" cy="16" r="2" fill="currentColor"/>`),
		g.Group(children),
	)
}