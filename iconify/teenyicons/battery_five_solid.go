package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryFiveSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 11.5A1.5 1.5 0 0 0 1.5 13h10a1.5 1.5 0 0 0 1.5-1.5v-8A1.5 1.5 0 0 0 11.5 2h-10A1.5 1.5 0 0 0 0 3.5v8Zm3-.5V4H2v7h1Zm2 0V4H4v7h1Zm2-7v7H6V4h1Zm2 7V4H8v7h1Zm2-7v7h-1V4h1Z" clip-rule="evenodd"/><path fill="currentColor" d="M15 5v5h-1V5h1Z"/>`),
		g.Group(children),
	)
}