package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForNiger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#fff" d="M0 23h64v18H0z"/><path fill="#e05206" d="M54 10H10C3.373 10 0 14.925 0 21v2h64v-2c0-6.075-3.373-11-10-11"/><path fill="#137a08" d="M0 43c0 6.075 3.373 11 10 11h44c6.627 0 10-4.925 10-11v-2H0v2"/><circle cx="31.979" cy="31.959" r="8.291" fill="#e05206"/>`),
		g.Group(children),
	)
}