package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoVideo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#4CAF50" d="M8 12h22c2.2 0 4 1.8 4 4v16c0 2.2-1.8 4-4 4H8c-2.2 0-4-1.8-4-4V16c0-2.2 1.8-4 4-4z"/><path fill="#388E3C" d="m44 35l-10-6V19l10-6z"/><path fill="none" stroke="#212121" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="4" d="m5 5l38 38"/>`),
		g.Group(children),
	)
}