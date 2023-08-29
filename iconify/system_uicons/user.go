package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func User(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" d="m11.5 4.5l2 2v1a3 3 0 0 1-5.995.176L7.5 6.5z"/><path d="M5.5 12V7.5a5 5 0 1 1 10 0V12"/><path stroke-linecap="round" d="M17.5 16.5v-.728c0-3.187-3.686-5.272-7-5.272s-7 2.085-7 5.272v.728a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1z"/></g>`),
		g.Group(children),
	)
}