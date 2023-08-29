package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HairCross(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 .4C4.697.4.399 4.698.399 10A9.6 9.6 0 0 0 10 19.601c5.301 0 9.6-4.298 9.6-9.601c0-5.302-4.299-9.6-9.6-9.6zm1 17.125V13H9v4.525A7.589 7.589 0 0 1 2.473 11H7V9H2.473A7.589 7.589 0 0 1 9 2.475V7h2V2.475A7.59 7.59 0 0 1 17.525 9H13v2h4.525A7.592 7.592 0 0 1 11 17.525z"/>`),
		g.Group(children),
	)
}