package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignDistributeTopSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1.5 1a.5.5 0 0 0 0 1h13a.5.5 0 0 0 0-1h-13Zm0 7a.5.5 0 0 0 0 1h13a.5.5 0 0 0 0-1h-13Zm8 6a1.5 1.5 0 0 0 1.5-1.5v-1A1.5 1.5 0 0 0 9.5 10h-3A1.5 1.5 0 0 0 5 11.5v1A1.5 1.5 0 0 0 6.5 14h3Zm2-7A1.5 1.5 0 0 0 13 5.5v-1A1.5 1.5 0 0 0 11.5 3h-7A1.5 1.5 0 0 0 3 4.5v1A1.5 1.5 0 0 0 4.5 7h7Z"/>`),
		g.Group(children),
	)
}