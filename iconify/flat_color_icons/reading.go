package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reading(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#5C6BC0" d="M40 40c-6.9 0-16 4-16 4V22s9-4 18-4l-2 22z"/><path fill="#7986CB" d="M8 40c6.9 0 16 4 16 4V22s-9-4-18-4l2 22z"/><g fill="#FFB74D"><circle cx="24" cy="12" r="8"/><path d="M41 32h1c.6 0 1-.4 1-1v-4c0-.6-.4-1-1-1h-1c-1.7 0-3 1.3-3 3s1.3 3 3 3zM7 26H6c-.6 0-1 .4-1 1v4c0 .6.4 1 1 1h1c1.7 0 3-1.3 3-3s-1.3-3-3-3z"/></g>`),
		g.Group(children),
	)
}