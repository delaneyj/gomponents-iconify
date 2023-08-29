package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScanPersonSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.5 2A2.5 2.5 0 0 0 2 4.5v1a.5.5 0 0 0 1 0v-1A1.5 1.5 0 0 1 4.5 3h1a.5.5 0 0 0 0-1h-1Zm6 0a.5.5 0 0 0 0 1h1A1.5 1.5 0 0 1 13 4.5v1a.5.5 0 0 0 1 0v-1A2.5 2.5 0 0 0 11.5 2h-1ZM3 10.5a.5.5 0 0 0-1 0v1A2.5 2.5 0 0 0 4.5 14h1a.5.5 0 0 0 0-1h-1A1.5 1.5 0 0 1 3 11.5v-1Zm11 0a.5.5 0 0 0-1 0v1a1.5 1.5 0 0 1-1.5 1.5h-1a.5.5 0 0 0 0 1h1a2.5 2.5 0 0 0 2.5-2.5v-1ZM9.085 13a1.5 1.5 0 0 1 1.415-1h1a.5.5 0 0 0 .5-.5a1.5 1.5 0 0 0-1.5-1.5h-5A1.5 1.5 0 0 0 4 11.5a.5.5 0 0 0 .5.5h1a1.5 1.5 0 0 1 1.415 1h2.17ZM8 9a2.25 2.25 0 1 0 0-4.5A2.25 2.25 0 0 0 8 9Z"/>`),
		g.Group(children),
	)
}