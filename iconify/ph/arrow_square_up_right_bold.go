package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSquareUpRightBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 28H48a20 20 0 0 0-20 20v160a20 20 0 0 0 20 20h160a20 20 0 0 0 20-20V48a20 20 0 0 0-20-20Zm-4 176H52V52h152ZM87.51 168.49a12 12 0 0 1 0-17L131 108h-19a12 12 0 0 1 0-24h48a12 12 0 0 1 12 12v48a12 12 0 0 1-24 0v-19l-43.51 43.52a12 12 0 0 1-17 0Z"/>`),
		g.Group(children),
	)
}