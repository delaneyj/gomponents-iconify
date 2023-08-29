package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareAsteriskFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M19 2a3 3 0 0 1 3 3v14a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V5a3 3 0 0 1 3-3zm-7 5.5a1 1 0 0 0-1 1v1.631l-1.445-.963l-.101-.06a1 1 0 0 0-1.009 1.724L10.195 12l-1.75 1.169l-.093.07a1 1 0 0 0 1.203 1.594L11 13.868V15.5l.007.117A1 1 0 0 0 13 15.5v-1.631l1.445.963l.101.06a1 1 0 0 0 1.009-1.724l-1.752-1.169l1.752-1.167l.093-.07a1 1 0 0 0-1.203-1.594L13 10.13V8.5l-.007-.117A1 1 0 0 0 12 7.5z"/></g>`),
		g.Group(children),
	)
}