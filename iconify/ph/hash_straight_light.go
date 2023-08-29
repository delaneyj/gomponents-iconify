package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HashStraightLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 154h-50v-52h50a6 6 0 0 0 0-12h-50V40a6 6 0 0 0-12 0v50h-52V40a6 6 0 0 0-12 0v50H40a6 6 0 0 0 0 12h50v52H40a6 6 0 0 0 0 12h50v50a6 6 0 0 0 12 0v-50h52v50a6 6 0 0 0 12 0v-50h50a6 6 0 0 0 0-12Zm-114 0v-52h52v52Z"/>`),
		g.Group(children),
	)
}