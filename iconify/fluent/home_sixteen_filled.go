package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.687 1.262a1 1 0 0 0-1.374 0L2.469 5.84A1.5 1.5 0 0 0 2 6.931v5.57A1.5 1.5 0 0 0 3.5 14H5a1.5 1.5 0 0 0 1.5-1.5V10a.5.5 0 0 1 .5-.5h2a.5.5 0 0 1 .5.5v2.5A1.5 1.5 0 0 0 11 14h1.5a1.5 1.5 0 0 0 1.5-1.5V6.93a1.5 1.5 0 0 0-.47-1.09L8.687 1.26Z"/>`),
		g.Group(children),
	)
}