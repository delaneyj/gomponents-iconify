package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwipeRightTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 12a1 1 0 0 0 1 1h11.585l-1.292 1.293a1 1 0 0 0-.083 1.32l.083.094a1 1 0 0 0 1.32.083l.094-.083l3-3a1 1 0 0 0 .083-1.32l-.083-.094l-3-3a1 1 0 0 0-1.497 1.32l.083.094L18.585 11H7a1 1 0 0 0-1 1Zm-4 0a5 5 0 0 0 9.584 2H9.872a3.5 3.5 0 1 1 0-4h1.712A5.001 5.001 0 0 0 2 12Z"/>`),
		g.Group(children),
	)
}