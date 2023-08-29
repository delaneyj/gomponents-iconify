package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Webcam(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M13 10a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/><path fill-rule="evenodd" d="M13 14.9A5.002 5.002 0 0 0 12 5a5 5 0 0 0-1 9.9V17H7v2h10v-2h-4v-2.1ZM12 13a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}