package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastDownButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="m55 31l-9.111 18.795l-7.668 15.817c-.897 1.85-3.545 1.85-4.442 0L26.11 49.795L17 31"/><path d="m55 5l-9.111 18.882l-7.668 15.891c-.897 1.86-3.545 1.86-4.442 0l-7.668-15.89L17 5"/></g>`),
		g.Group(children),
	)
}