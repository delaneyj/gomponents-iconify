package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blitzerde(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 20.35L7.75 30.44L4.48 42.49L24 28.72l19.5 13.77l-3.27-12Zm0-8.58L11.13 18L9 25.59l15-8.73l15 8.73L36.87 18Zm0-6.27L13.34 9.9L11.9 15L24 9.27L36.1 15l-1.44-5.1Z"/>`),
		g.Group(children),
	)
}