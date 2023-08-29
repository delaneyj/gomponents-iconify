package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DustTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.75 3.5a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Zm2.34 2.652A4.001 4.001 0 1 1 14.646 10a4 4 0 1 1-6.556 3.85a4 4 0 1 1 0-7.699ZM5.75 18a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5ZM19 10a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Z"/>`),
		g.Group(children),
	)
}