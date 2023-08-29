package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdobePhotoshopMix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="21" cy="21" r="7.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="26.5" cy="26.5" r="7.5" fill="none" stroke="currentColor" stroke-dasharray="0 0 1 1.78" stroke-dashoffset="1.6" stroke-linecap="round" stroke-linejoin="round"/><rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4" ry="4"/>`),
		g.Group(children),
	)
}