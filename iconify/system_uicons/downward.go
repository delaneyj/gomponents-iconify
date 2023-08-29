package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Downward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.23 10.23c3.264-3.79 6.687-5.033 10.27-3.73c-3.552.646-6.009 2.855-7.371 6.63L12.5 15.5h-8v-8z"/>`),
		g.Group(children),
	)
}