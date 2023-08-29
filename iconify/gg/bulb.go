package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bulb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M4 9a7.997 7.997 0 0 0 4 6.93V16a4 4 0 1 0 8 0v-.07A8 8 0 1 0 4 9Zm12 4.472a6 6 0 1 0-8 0h2V16a2 2 0 1 0 4 0v-2.53l2 .001Z" clip-rule="evenodd"/><path d="M10 21.006V21c.588.34 1.271.535 2 .535c.729 0 1.412-.195 2-.535v.006a2 2 0 1 1-4 0Z"/></g>`),
		g.Group(children),
	)
}