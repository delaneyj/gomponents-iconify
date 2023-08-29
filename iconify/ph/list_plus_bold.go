package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListPlusBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M28 64a12 12 0 0 1 12-12h176a12 12 0 0 1 0 24H40a12 12 0 0 1-12-12Zm12 76h176a12 12 0 0 0 0-24H40a12 12 0 0 0 0 24Zm104 40H40a12 12 0 0 0 0 24h104a12 12 0 0 0 0-24Zm88 0h-12v-12a12 12 0 0 0-24 0v12h-12a12 12 0 0 0 0 24h12v12a12 12 0 0 0 24 0v-12h12a12 12 0 0 0 0-24Z"/>`),
		g.Group(children),
	)
}