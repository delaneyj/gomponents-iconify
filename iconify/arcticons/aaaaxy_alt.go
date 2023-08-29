package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AaaaxyAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 5.5v31.559l5.441 5.441H37.06l5.441-5.441V5.5h-37Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.941 17.47v13.06l6.53 6.529h13.058l6.53-6.53V17.471l-6.53-6.53H17.471l-6.53 6.53Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.382 22.912L24 15.294l7.618 7.618M24 15.294v17.412"/>`),
		g.Group(children),
	)
}