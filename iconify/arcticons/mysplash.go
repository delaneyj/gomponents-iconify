package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mysplash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.5 35.29L24 24l-6.5 11.29h-13l13-22.58L24 24l6.5-11.29l13 22.58Z"/>`),
		g.Group(children),
	)
}