package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Math(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="23" x="5.5" y="12.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><rect width="37" height="23" x="5.5" y="12.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" transform="rotate(90 24 24)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.13 21.97l-4.94 6.55m4.94 0l-4.94-6.55M31.45 19h-8.11l-3.88 10l-2.91-7.03"/>`),
		g.Group(children),
	)
}