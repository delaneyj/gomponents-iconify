package wi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoonWaningCrescentFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 30 30"),
		g.Raw(`<path fill="currentColor" d="M3.74 14.44c0 2.04.5 3.93 1.51 5.65s2.37 3.1 4.1 4.1S12.96 25.7 15 25.7c-2.07-1.01-3.59-2.45-4.56-4.33S9 17.19 9 14.44c0-2.53.56-4.78 1.69-6.75s2.57-3.47 4.31-4.5c-2.04 0-3.93.5-5.65 1.51s-3.1 2.37-4.1 4.09s-1.51 3.61-1.51 5.65z"/>`),
		g.Group(children),
	)
}