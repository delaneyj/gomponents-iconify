package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxOuterLightRoundUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M22 22h-2v-5h-1v-4h-1v-2h-1V9h-1V8h-1V7h-1V6h-1V5h-2V4H9V3H5V2H.004V0H6v1h4v1h2v1h2v1h1v1h1v1h1v1h1v1h1v2h1v2h1v4h1v6Z"/>`),
		g.Group(children),
	)
}