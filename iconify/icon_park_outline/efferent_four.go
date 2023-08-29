package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EfferentFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20 5H10a2 2 0 0 0-2 2v34a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V24.75M28 5h12v12m-19 7L39 6"/>`),
		g.Group(children),
	)
}