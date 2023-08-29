package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForSuriname(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#137a08" d="M0 43c0 6.075 3.373 11 10 11h44c6.627 0 10-4.925 10-11v-3H0v3m54-33H10C3.373 10 0 14.925 0 21v3h64v-3c0-6.075-3.373-11-10-11"/><path fill="#ec1c24" d="M0 24h64v16H0z"/><path fill="#e6e7e8" d="M0 21h64v3H0zm0 19h64v3H0z"/><path fill="#f9cb38" d="m39.39 30.15l-5.689.01l-1.763-5.755l-1.748 5.755l-5.692-.01l4.612 3.496l-1.79 5.714l4.627-3.55l4.627 3.55l-1.79-5.714z"/>`),
		g.Group(children),
	)
}