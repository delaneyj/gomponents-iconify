package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spanish(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 37.5v-27c0-1.1-.9-2-2-2h-35c-1.1 0-2 .9-2 2v27c0 1.1.9 2 2 2h35c1.1 0 2-.9 2-2Zm0-6.5h-39m0-14h39"/>`),
		g.Group(children),
	)
}