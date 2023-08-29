package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Strong(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.5 24h5m-28 0h15m-28 0h5"/><rect width="8" height="17" x="8.5" y="15.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1" ry="1"/><rect width="8" height="17" x="31.5" y="15.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1" ry="1"/>`),
		g.Group(children),
	)
}