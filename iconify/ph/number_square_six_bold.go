package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSquareSixBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 28H48a20 20 0 0 0-20 20v160a20 20 0 0 0 20 20h160a20 20 0 0 0 20-20V48a20 20 0 0 0-20-20Zm-4 176H52V52h152Zm-76-16a40 40 0 0 0 5.06-79.67l13.24-22.18a12 12 0 1 0-20.6-12.3l-32.24 54A40 40 0 0 0 128 188Zm0-56a16 16 0 1 1-16 16a16 16 0 0 1 16-16Z"/>`),
		g.Group(children),
	)
}