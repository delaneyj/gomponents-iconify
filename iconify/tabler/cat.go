package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 3v10a8 8 0 1 1-16 0V3l3.432 3.432A7.963 7.963 0 0 1 12 5c1.769 0 3.403.574 4.728 1.546L20 3z"/><path d="M2 16h5l-4 4m19-4h-5l4 4m-10-4a1 1 0 1 0 2 0a1 1 0 1 0-2 0m-2-5v.01m6-.01v.01"/></g>`),
		g.Group(children),
	)
}