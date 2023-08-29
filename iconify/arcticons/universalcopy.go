package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Universalcopy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.375 19.788h19.656m-19.656 5.989h11.832m2.608 3.189h2.67a2.67 2.67 0 0 1 2.67 2.67v0a2.67 2.67 0 0 1-2.67 2.67h0a2.67 2.67 0 0 1-2.67-2.67v-2.67h0Zm-16.951-6.294v2.67a2.67 2.67 0 0 1-2.67 2.67h0a2.67 2.67 0 0 1-2.67-2.67h0a2.67 2.67 0 0 1 2.67-2.67h2.67Z"/>`),
		g.Group(children),
	)
}