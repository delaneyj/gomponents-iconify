package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Template(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#BBDEFB" d="M4 7h40v34H4z"/><path fill="#3F51B5" d="M9 12h30v5H9z"/><path fill="#2196F3" d="M9 21h13v16H9zm17 0h13v16H26z"/>`),
		g.Group(children),
	)
}