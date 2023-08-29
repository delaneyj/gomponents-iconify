package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusinessContact(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#673AB7" d="M40 7H8c-2.2 0-4 1.8-4 4v26c0 2.2 1.8 4 4 4h5v-1.3c-.6-.3-1-1-1-1.7c0-1.1.9-2 2-2s2 .9 2 2c0 .7-.4 1.4-1 1.7V41h18v-1.3c-.6-.3-1-1-1-1.7c0-1.1.9-2 2-2s2 .9 2 2c0 .7-.4 1.4-1 1.7V41h5c2.2 0 4-1.8 4-4V11c0-2.2-1.8-4-4-4z"/><g fill="#D1C4E9"><circle cx="24" cy="18" r="4"/><path d="M31 28s-1.9-4-7-4s-7 4-7 4v2h14v-2z"/></g>`),
		g.Group(children),
	)
}