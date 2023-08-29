package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapArrowDownBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M8.037 14.144a.5.5 0 0 1-.68-.244L3.164 4.498c-.667-1.495.815-3.047 2.202-2.306l5.904 3.152c.46.245 1 .245 1.459 0l5.904-3.152c1.387-.741 2.869.81 2.202 2.306l-1.572 3.525a2 2 0 0 1-.932.974L8.037 14.144Z"/><path d="M8.61 15.534a.5.5 0 0 0-.234.651l2.151 4.823c.59 1.323 2.355 1.323 2.945 0l3.968-8.898a.5.5 0 0 0-.68-.651l-8.15 4.075Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}