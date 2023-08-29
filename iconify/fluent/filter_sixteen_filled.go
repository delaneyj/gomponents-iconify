package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1.75 3.75A.75.75 0 0 1 2.5 3h11a.75.75 0 0 1 0 1.5h-11a.75.75 0 0 1-.75-.75Zm2 4A.75.75 0 0 1 4.5 7h7a.75.75 0 0 1 0 1.5h-7a.75.75 0 0 1-.75-.75Zm2 4A.75.75 0 0 1 6.5 11h3a.75.75 0 0 1 0 1.5h-3a.75.75 0 0 1-.75-.75Z"/>`),
		g.Group(children),
	)
}