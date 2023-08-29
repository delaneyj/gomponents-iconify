package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyBtcLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M166.69 116.41A42 42 0 0 0 142 42.05V24a6 6 0 0 0-12 0v18h-20V24a6 6 0 0 0-12 0v18H64a6 6 0 0 0 0 12h10v140H64a6 6 0 0 0 0 12h34v18a6 6 0 0 0 12 0v-18h20v18a6 6 0 0 0 12 0v-18h10a46 46 0 0 0 14.69-89.59ZM170 84a30 30 0 0 1-30 30H86V54h54a30 30 0 0 1 30 30Zm-18 110H86v-68h66a34 34 0 0 1 0 68Z"/>`),
		g.Group(children),
	)
}