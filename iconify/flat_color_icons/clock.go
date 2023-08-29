package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="20" fill="#00ACC1"/><circle cx="24" cy="24" r="16" fill="#eee"/><path d="M23 11h2v13h-2z"/><path d="M31.285 29.654L29.66 31.28l-6.504-6.504l1.626-1.627z"/><circle cx="24" cy="24" r="2"/><circle cx="24" cy="24" r="1" fill="#00ACC1"/>`),
		g.Group(children),
	)
}