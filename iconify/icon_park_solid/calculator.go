package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calculator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCalculator0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M40 4H8a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2h32a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2Z"/><path fill="#000" stroke="#000" d="M35 10H13v9h22v-9Z"/><path stroke="#000" stroke-linecap="round" d="m12 28l7 7m0-7l-7 7m16 0h8m-8-6h8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCalculator0)"/>`),
		g.Group(children),
	)
}