package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlashForwardSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.267 2.05a.75.75 0 0 1 .434.967l-4 10.5a.75.75 0 0 1-1.402-.534l4-10.5a.75.75 0 0 1 .968-.434Z"/>`),
		g.Group(children),
	)
}