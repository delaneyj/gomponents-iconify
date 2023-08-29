package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSquareSevenBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 28H48a20 20 0 0 0-20 20v160a20 20 0 0 0 20 20h160a20 20 0 0 0 20-20V48a20 20 0 0 0-20-20Zm-4 176H52V52h152ZM92 88a12 12 0 0 1 12-12h48a12 12 0 0 1 11.28 16.1l-32 88a12 12 0 0 1-22.56-8.2l26.15-71.9H104a12 12 0 0 1-12-12Z"/>`),
		g.Group(children),
	)
}