package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandDebian(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 17c-2.397-.943-4-3.153-4-5.635c0-2.19 1.039-3.14 1.604-3.595C12.25 5.637 16 7.5 16 11c0 2.5-2.905 2.121-3.5 1.5c-.595-.621-1-1.5-.5-2.5"/><path d="M3 12a9 9 0 1 0 18 0a9 9 0 1 0-18 0"/></g>`),
		g.Group(children),
	)
}