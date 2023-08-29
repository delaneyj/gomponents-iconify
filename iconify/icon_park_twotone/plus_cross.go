package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlusCross(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPlusCross0"><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M30 4H18v14H4v12h14v14h12V30h14V18H30V4Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPlusCross0)"/>`),
		g.Group(children),
	)
}