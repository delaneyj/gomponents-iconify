package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Monitoring(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 18h2v3h-2zm5 3v2H8v-2zm4-3H4a3.003 3.003 0 0 1-3-3V4a3.003 3.003 0 0 1 3-3h16a3.003 3.003 0 0 1 3 3v11a3.003 3.003 0 0 1-3 3ZM4 3a1 1 0 0 0-1 1v11a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1Z"/><path fill="currentColor" d="m16 15l-1.914-6.38L13 13l-1.309-3h-.331L10 14L8.843 9.933L8.309 11H5v-1h2.691L9 7l1.068 3.713L10.64 9h1.669l.487.973L14 4l2 8l.64-2H19v1h-1.64L16 15z"/>`),
		g.Group(children),
	)
}