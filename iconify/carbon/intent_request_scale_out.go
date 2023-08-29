package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IntentRequestScaleOut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m19 20.4l1.4-1.4l7.6 7.6V20h2v10H20v-2h6.6zm-6 0L11.6 19L4 26.6V20H2v10h10v-2H5.4zm4-4.4h-2V5.8l-4.6 4.6L9 9l7-7l7 7l-1.4 1.4L17 5.8z"/>`),
		g.Group(children),
	)
}