package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M513 216.6H158.5L316.1 59.1H197.9L1 256l196.9 196.9h118.2L158.5 295.4H513z"/>`),
		g.Group(children),
	)
}