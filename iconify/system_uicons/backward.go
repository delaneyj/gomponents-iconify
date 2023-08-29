package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Backward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5 7.5c5.185-.47 8.52 1.53 10 6c-2.825-3.14-6.341-3.718-10-2v3l-5-5l5-5z"/>`),
		g.Group(children),
	)
}