package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForGambia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#002664" d="M0 25h64v14H0z"/><path fill="#e6e7e8" d="M0 24h64v3H0zm0 13h64v3H0z"/><path fill="#ec1c24" d="M54 10H10C3.373 10 0 14.925 0 21v3h64v-3c0-6.075-3.373-11-10-11"/><path fill="#137a08" d="M0 43c0 6.075 3.373 11 10 11h44c6.627 0 10-4.925 10-11v-3H0v3"/>`),
		g.Group(children),
	)
}