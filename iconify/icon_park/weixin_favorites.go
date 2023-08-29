package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeixinFavorites(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 14L24 6L42 14M6 14L24 22M6 14V34L24 42M42 14L24 22M42 14V34L24 42M24 22V42"/>`),
		g.Group(children),
	)
}