package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SerialTasks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#90CAF9" d="M33 9H11v4h22c1.1 0 2 .9 2 2v20H23v4h16V15c0-3.3-2.7-6-6-6z"/><path fill="#D81B60" d="M6 6h10v10H6z"/><g fill="#2196F3"><path d="M32 17h10v10H32zM16 32h10v10H16z"/><circle cx="26" cy="11" r="5"/><circle cx="37" cy="37" r="5"/></g>`),
		g.Group(children),
	)
}