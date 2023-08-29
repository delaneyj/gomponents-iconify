package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TownHall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M13 4H9V1L7.5 0L6 1v3H2L1 5v1h13V5l-1-1zM7.5 1.5c.4 0 .7.3.7.8s-.3.7-.7.7s-.8-.3-.8-.8c0-.4.4-.7.8-.7zM13 7H2v4l-1 1.5V14h13v-1.5L13 11V7zm-8 5.5H4V8h1v4.5zm3 0H7V8h1v4.5zm3 0h-1V8h1v4.5z"/>`),
		g.Group(children),
	)
}