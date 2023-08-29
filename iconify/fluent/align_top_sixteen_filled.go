package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignTopSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1.5 2a.5.5 0 0 0 0 1h13a.5.5 0 0 0 0-1h-13ZM2 5.75C2 4.784 2.784 4 3.75 4h1.5C6.216 4 7 4.784 7 5.75v6.5A1.75 1.75 0 0 1 5.25 14h-1.5A1.75 1.75 0 0 1 2 12.25v-6.5Zm7 0C9 4.784 9.784 4 10.75 4h1.5c.966 0 1.75.784 1.75 1.75v4.5A1.75 1.75 0 0 1 12.25 12h-1.5A1.75 1.75 0 0 1 9 10.25v-4.5Z"/>`),
		g.Group(children),
	)
}