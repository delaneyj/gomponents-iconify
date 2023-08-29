package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NdsFourDroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="16.4" height="12.2" x="15.8" y="10.3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1" ry="1"/><rect width="16.4" height="12.2" x="15.8" y="25.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1" ry="1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v31l4 4h31a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}