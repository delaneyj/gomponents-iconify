package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CapsuleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="m12 2l-.243.004A7.004 7.004 0 0 0 5 9v6a7 7 0 0 0 7 7l.243-.004A7.004 7.004 0 0 0 19 15V9a7 7 0 0 0-7-7z"/></g>`),
		g.Group(children),
	)
}