package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSquareFiveBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 28H48a20 20 0 0 0-20 20v160a20 20 0 0 0 20 20h160a20 20 0 0 0 20-20V48a20 20 0 0 0-20-20Zm-4 176H52V52h152ZM92.16 126l8-48A12 12 0 0 1 112 68h40a12 12 0 0 1 0 24h-29.83l-2.71 16.23A45.39 45.39 0 0 1 124 108a40 40 0 0 1 0 80a39.53 39.53 0 0 1-28.57-11.6a12 12 0 1 1 17.14-16.8A15.54 15.54 0 0 0 124 164a16 16 0 0 0 0-32a15.54 15.54 0 0 0-11.43 4.4A12 12 0 0 1 92.16 126Z"/>`),
		g.Group(children),
	)
}