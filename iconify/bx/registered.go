package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Registered(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.14 2a10 10 0 1 0 10 10a10 10 0 0 0-10-10zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8z"/><path fill="currentColor" d="M16.14 10a3 3 0 0 0-3-3h-5v10h2v-4h1.46l2.67 4h2.4l-2.75-4.12A3 3 0 0 0 16.14 10zm-3 1h-3V9h3a1 1 0 0 1 0 2z"/>`),
		g.Group(children),
	)
}