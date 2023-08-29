package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calculator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCalculator0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M40 4H8a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2h32a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2Z"/><path fill="#555" d="M35 10H13v9h22v-9Z"/><path stroke-linecap="round" d="m12 28l7 7m0-7l-7 7m16 0h8m-8-6h8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCalculator0)"/>`),
		g.Group(children),
	)
}