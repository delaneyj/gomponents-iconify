package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSquareThreeBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 28H48a20 20 0 0 0-20 20v160a20 20 0 0 0 20 20h160a20 20 0 0 0 20-20V48a20 20 0 0 0-20-20Zm-4 176H52V52h152ZM92 80a12 12 0 0 1 12-12h48a12 12 0 0 1 9.83 18.88l-18.34 26.2A40 40 0 1 1 95.43 176a12 12 0 1 1 17.14-16.79A16 16 0 1 0 124 132a12 12 0 0 1-9.83-18.88L129 92h-25a12 12 0 0 1-12-12Z"/>`),
		g.Group(children),
	)
}