package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Picture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#F57C00" d="M40 41H8c-2.2 0-4-1.8-4-4V11c0-2.2 1.8-4 4-4h32c2.2 0 4 1.8 4 4v26c0 2.2-1.8 4-4 4z"/><circle cx="35" cy="16" r="3" fill="#FFF9C4"/><path fill="#942A09" d="M20 16L9 32h22z"/><path fill="#BF360C" d="m31 22l-8 10h16z"/>`),
		g.Group(children),
	)
}