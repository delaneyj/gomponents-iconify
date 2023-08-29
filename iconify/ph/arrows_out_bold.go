package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsOutBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M220 48v48a12 12 0 0 1-24 0V77l-35.51 35.52a12 12 0 0 1-17-17L179 60h-19a12 12 0 0 1 0-24h48a12 12 0 0 1 12 12ZM95.51 143.51L60 179v-19a12 12 0 0 0-24 0v48a12 12 0 0 0 12 12h48a12 12 0 0 0 0-24H77l35.52-35.51a12 12 0 0 0-17-17ZM208 148a12 12 0 0 0-12 12v19l-35.51-35.52a12 12 0 0 0-17 17L179 196h-19a12 12 0 0 0 0 24h48a12 12 0 0 0 12-12v-48a12 12 0 0 0-12-12ZM77 60h19a12 12 0 0 0 0-24H48a12 12 0 0 0-12 12v48a12 12 0 0 0 24 0V77l35.51 35.52a12 12 0 0 0 17-17Z"/>`),
		g.Group(children),
	)
}