package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M11 1V0H5v1H3v1H2v14h12v-1h1V1h-4zM6 1h4v2H6V1zm7 14H3V3h2v1h6V3h2v12z"/><path fill="currentColor" d="m7.39 12.47l-3-2.73l1.35-1.48L7.32 9.7l2.87-2.9l1.42 1.4l-4.22 4.27z"/>`),
		g.Group(children),
	)
}