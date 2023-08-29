package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scanner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M358 164q11 3 18.5 14.5T384 203v117q0 18-12.5 30.5T341 363H43q-18 0-30.5-12.5T0 320v-85q0-18 12.5-30.5T43 192h268L11 83l15-40zM85 299v-43H43v43h42zm256 0v-43H128v43h213z"/>`),
		g.Group(children),
	)
}