package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirplayAudio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m12 16l6 5H6l6-5Z"/><path d="M4 18a9.956 9.956 0 0 1-2-6C2 6.477 6.477 2 12 2s10 4.477 10 10a9.956 9.956 0 0 1-2 6"/><path d="M17.123 15.125a6 6 0 1 0-10.247-.002"/><path d="M14 12a2 2 0 1 0-4 0"/></g>`),
		g.Group(children),
	)
}