package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSquareUpLeftBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 28H48a20 20 0 0 0-20 20v160a20 20 0 0 0 20 20h160a20 20 0 0 0 20-20V48a20 20 0 0 0-20-20Zm-4 176H52V52h152ZM84 144V96a12 12 0 0 1 12-12h48a12 12 0 0 1 0 24h-19l43.52 43.51a12 12 0 0 1-17 17L108 125v19a12 12 0 0 1-24 0Z"/>`),
		g.Group(children),
	)
}