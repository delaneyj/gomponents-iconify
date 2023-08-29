package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Way(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="32.356" height="23.23" x="11.144" y="9.063" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.898"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.144 15.707H7.398A2.897 2.897 0 0 0 4.5 18.605v17.434a2.898 2.898 0 0 0 2.898 2.898h26.56a2.898 2.898 0 0 0 2.898-2.898v-3.746M11.144 12.809H43.5m-32.356 5.796H43.5"/>`),
		g.Group(children),
	)
}