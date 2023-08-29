package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scribd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="15.895" cy="20.232" r="6.068" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.895 26.3c-9.78 0-8.249-14.316-1.473-19.265s15.023-1.885 21.68 2.651L33.57 13.28"/><circle cx="32.105" cy="27.763" r="6.068" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.105 21.695c9.78 0 8.249 14.316 1.473 19.265s-18.214 2.167-24.214-4.055l2.533-3.594"/>`),
		g.Group(children),
	)
}