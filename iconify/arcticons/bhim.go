package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bhim(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.801 22.51a1.284 1.284 0 0 1-.294 1.592m.294-1.592L21.171 4.5L10.19 42.346l20.316-18.244m7.146-.451a1.284 1.284 0 0 1-.294 1.592m.294-1.592l-9.51-17.97l-2.219 7.424m11.435 12.138L17.194 43.5l2.884-9.93"/>`),
		g.Group(children),
	)
}