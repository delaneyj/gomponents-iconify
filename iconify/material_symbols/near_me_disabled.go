package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NearMeDisabled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12.9 21l-2.85-7.05L3 11.1V9.7l4.875-1.825L2.8 2.8l1.425-1.425l18.4 18.4L21.2 21.2l-5.075-5.075L14.3 21h-1.4Zm4.775-9.025l-5.65-5.65L21 3l-3.325 8.975Z"/>`),
		g.Group(children),
	)
}