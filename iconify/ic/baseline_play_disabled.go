package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaselinePlayDisabled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 5.19V5l11 7l-2.55 1.63L8 5.19zm12 14.54l-5.11-5.11L8 7.73L4.27 4L3 5.27l5 5V19l5.33-3.4l5.4 5.4L20 19.73z"/>`),
		g.Group(children),
	)
}