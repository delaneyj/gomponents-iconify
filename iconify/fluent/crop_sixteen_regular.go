package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 1.5a.5.5 0 0 0-1 0V4H1.5a.5.5 0 0 0 0 1H4v4.5A2.5 2.5 0 0 0 6.5 12H11v2.5a.5.5 0 0 0 1 0V12h2.5a.5.5 0 0 0 0-1h-8A1.5 1.5 0 0 1 5 9.5v-8Zm6 5V10h1V6.5A2.5 2.5 0 0 0 9.5 4H6v1h3.5A1.5 1.5 0 0 1 11 6.5Z"/>`),
		g.Group(children),
	)
}