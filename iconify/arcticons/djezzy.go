package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Djezzy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.242 4.711A1.558 1.558 0 0 0 7.906 6.06v35.88a1.558 1.558 0 0 0 2.336 1.349l31.073-17.94a1.558 1.558 0 0 0 0-2.698Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.694 30.5v-13h2.925a5.687 5.687 0 0 1 5.687 5.688v1.625A5.687 5.687 0 0 1 19.62 30.5Z"/>`),
		g.Group(children),
	)
}