package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLineDownLeftBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M228 40a12 12 0 0 1-12 12H40a12 12 0 0 1 0-24h176a12 12 0 0 1 12 12Zm-60.49 39.51L76 171v-67a12 12 0 0 0-24 0v96a12 12 0 0 0 12 12h96a12 12 0 0 0 0-24H93l91.52-91.51a12 12 0 0 0-17-17Z"/>`),
		g.Group(children),
	)
}