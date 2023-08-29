package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForBotswana(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#25333a" d="M0 25h64v14H0z"/><path fill="#e6e7e8" d="M0 23h64v4H0zm0 14h64v4H0z"/><path fill="#75aadb" d="M54 10H10C3.373 10 0 14.925 0 21v2h64v-2c0-6.075-3.373-11-10-11M0 43c0 6.075 3.373 11 10 11h44c6.627 0 10-4.925 10-11v-2H0v2"/>`),
		g.Group(children),
	)
}