package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileEight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M300.7 0H66v512h384V149.3H300.7V0zm10.6 245.3L354 288l-53.3 53.3l53.3 53.3l-42.7 42.7L258 384l-53.3 53.3l-42.7-42.6l53.3-53.3L162 288l42.7-42.7l53.3 53.3l53.3-53.3zM322 0v128h128L322 0z"/>`),
		g.Group(children),
	)
}