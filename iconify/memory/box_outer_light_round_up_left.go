package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxOuterLightRoundUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M22 0v2h-5v1h-4v1h-2v1H9v1H8v1H7v1H6v1H5v2H4v2H3v4H2v5H.004v-6H1v-4h1v-2h1V8h1V7h1V6h1V5h1V4h1V3h2V2h2V1h4V0h6Z"/>`),
		g.Group(children),
	)
}