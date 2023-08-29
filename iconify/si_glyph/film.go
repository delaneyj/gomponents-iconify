package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Film(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15 1h-.988v1h-2.017V1H4.101v1h-2.2V1H1v16h.901v-1h2.053v1h8.041v-1h2.017v1H15V1zM4 14H2v-2h2v2zm0-4H2V8h2v2zm0-4H2V4h2v2zm7 9H5v-5h6v5zm0-7H5V3h6v5zm3 6h-2v-2h2v2zm0-4h-2V8h2v2zm-2-4V4h2v2h-2z"/>`),
		g.Group(children),
	)
}