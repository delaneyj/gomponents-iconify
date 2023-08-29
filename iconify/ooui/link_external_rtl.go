package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkExternalRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 12H1v5c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V3c0-1.1-.9-2-2-2h-5v2h5v14H3z"/><path fill="currentColor" d="m1 9l3.3-3.3l5.7 5.7l1.4-1.4l-5.7-5.7L9 1H1z"/>`),
		g.Group(children),
	)
}