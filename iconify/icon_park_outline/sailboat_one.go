package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SailboatOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M21 31V5L7 31h14Zm6 0V13l14 18H27ZM5 37h38l-5 6H10l-5-6Z"/>`),
		g.Group(children),
	)
}