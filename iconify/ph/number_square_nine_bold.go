package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSquareNineBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 28H48a20 20 0 0 0-20 20v160a20 20 0 0 0 20 20h160a20 20 0 0 0 20-20V48a20 20 0 0 0-20-20Zm-4 176H52V52h152Zm-96-61.36a39.81 39.81 0 0 0 14.94 5l-13.24 22.21a12 12 0 1 0 20.6 12.3L162.64 128A40 40 0 1 0 108 142.64Zm6.14-42.64a16 16 0 0 1 27.72 16a16 16 0 0 1-27.7-16Z"/>`),
		g.Group(children),
	)
}