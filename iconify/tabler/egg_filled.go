package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EggFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M12.002 2C7.829 1.992 4 8.058 4 14.083C4 18.791 7.25 22 12 22c4.727-.206 8-3.328 8-7.917C20 8.063 16.175 2.008 12.002 2z"/></g>`),
		g.Group(children),
	)
}