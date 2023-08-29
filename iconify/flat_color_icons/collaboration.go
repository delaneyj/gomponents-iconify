package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Collaboration(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#1565C0" d="M25 22h13l6 6V11c0-2.2-1.8-4-4-4H25c-2.2 0-4 1.8-4 4v7c0 2.2 1.8 4 4 4z"/><path fill="#2196F3" d="M23 19H10l-6 6V8c0-2.2 1.8-4 4-4h15c2.2 0 4 1.8 4 4v7c0 2.2-1.8 4-4 4z"/><g fill="#FFA726"><circle cx="12" cy="31" r="5"/><circle cx="36" cy="31" r="5"/></g><path fill="#607D8B" d="M20 42s-2.2-4-8-4s-8 4-8 4v2h16v-2zm24 0s-2.2-4-8-4s-8 4-8 4v2h16v-2z"/>`),
		g.Group(children),
	)
}