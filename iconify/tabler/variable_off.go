package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VariableOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4.675 4.68C2.505 9.456 2.613 14.272 5 20M19 4c1.959 3.917 2.383 7.834 1.272 12.232m-.983 3.051c-.093.238-.189.477-.289.717m-7.304-8.304c.095.257.2.533.32.831C13 15 13 16 14 16h1"/><path d="M8 16c1.5 0 3-2 4-3.5m2.022-2.514C14.651 9.404 15.326 9 16 9M3 3l18 18"/></g>`),
		g.Group(children),
	)
}