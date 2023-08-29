package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoinFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23.005 12.003v2c0 3.314-4.925 6-11 6c-5.967 0-10.824-2.591-10.995-5.823l-.005-.177v-2c0 3.314 4.925 6 11 6s11-2.686 11-6Zm-11-8c6.075 0 11 2.686 11 6s-4.925 6-11 6s-11-2.686-11-6s4.925-6 11-6Z"/>`),
		g.Group(children),
	)
}