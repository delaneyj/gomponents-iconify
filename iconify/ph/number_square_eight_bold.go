package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSquareEightBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 28H48a20 20 0 0 0-20 20v160a20 20 0 0 0 20 20h160a20 20 0 0 0 20-20V48a20 20 0 0 0-20-20Zm-4 176H52V52h152ZM88 152a40 40 0 1 0 67.6-28.91a36 36 0 1 0-55.2 0A39.87 39.87 0 0 0 88 152Zm40 16a16 16 0 1 1 16-16a16 16 0 0 1-16 16Zm-12-68a12 12 0 1 1 12 12a12 12 0 0 1-12-12Z"/>`),
		g.Group(children),
	)
}