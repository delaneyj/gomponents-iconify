package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IntentRequestScaleIn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m28.6 30l1.4-1.4l-7.6-7.6H29v-2H19v10h2v-6.6zM2 28.6L3.4 30l7.6-7.6V29h2V19H3v2h6.6zM17 2h-2v10.2l-4.6-4.6L9 9l7 7l7-7l-1.4-1.4l-4.6 4.6z"/>`),
		g.Group(children),
	)
}