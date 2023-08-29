package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Building(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M3 0v16h4v-3h2v3h4V0H3zm3 12H4v-2h2v2zm0-3H4V7h2v2zm0-3H4V4h2v2zm0-3H4V1h2v2zm3 9H7v-2h2v2zm0-3H7V7h2v2zm0-3H7V4h2v2zm0-3H7V1h2v2zm3 9h-2v-2h2v2zm0-3h-2V7h2v2zm0-3h-2V4h2v2zm0-3h-2V1h2v2z"/>`),
		g.Group(children),
	)
}