package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayCircleO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0Zm0 1.395a8.605 8.605 0 1 0 0 17.21a8.605 8.605 0 0 0 0-17.21ZM8.182 5.904c0-.468.436-.586.718-.281l3.675 3.962a.644.644 0 0 1 0 .831L8.9 14.377c-.282.305-.718.187-.718-.28Z"/>`),
		g.Group(children),
	)
}