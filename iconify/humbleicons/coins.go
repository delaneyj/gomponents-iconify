package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coins(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16.5 9.09a5.502 5.502 0 0 1-1 10.91a5.5 5.5 0 0 1-4.9-3M7.5 8l1-.5v4m5.5-2a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0Z"/>`),
		g.Group(children),
	)
}