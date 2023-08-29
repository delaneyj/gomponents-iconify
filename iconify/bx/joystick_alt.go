package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JoystickAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="15" cy="13" r="1" fill="currentColor"/><circle cx="17" cy="11" r="1" fill="currentColor"/><path fill="currentColor" d="M10 9H8v2H6v2h2v2h2v-2h2v-2h-2z"/><path fill="currentColor" d="M15 5H9a7 7 0 0 0-7 7a7 7 0 0 0 7 7h6a7 7 0 0 0 7-7a7 7 0 0 0-7-7zm0 12H9A5 5 0 1 1 9 7h6a5 5 0 1 1 0 10z"/>`),
		g.Group(children),
	)
}