package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crib(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M38 6v38m-28-4h28m-28-6h28M6 12h36m-7-6h6M7 6h6m-3 0v38m7-32v16m7-16v16m7-16v16"/>`),
		g.Group(children),
	)
}