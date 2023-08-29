package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForArmenia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#0033a0" d="M0 26h64v14H0z"/><path fill="#d90012" d="M54 11H10C3.373 11 0 15.925 0 22v4h64v-4c0-6.075-3.373-11-10-11"/><path fill="#f2a800" d="M0 44c0 6.075 3.373 11 10 11h44c6.627 0 10-4.925 10-11v-4H0v4"/>`),
		g.Group(children),
	)
}