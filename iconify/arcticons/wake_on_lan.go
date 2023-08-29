package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WakeOnLan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M28.662 11.167a16.503 16.503 0 1 1-9.016-.088"/><path d="M24.004 4.5h.096v25.18h-.096z"/></g>`),
		g.Group(children),
	)
}