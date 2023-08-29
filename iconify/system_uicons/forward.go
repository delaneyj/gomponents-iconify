package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.5 7.5c-5.185-.47-8.52 1.53-10 6c2.825-3.14 6.341-3.718 10-2v3l5-5l-5-5z"/>`),
		g.Group(children),
	)
}