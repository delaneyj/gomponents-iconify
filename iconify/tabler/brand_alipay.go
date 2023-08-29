package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandAlipay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2zM7 7h10m-5-4v7"/><path d="M21 17.314C18.029 15.391 6 8.535 6 15.45C6 17.166 7.52 18 8.985 18c3.512 0 6.814-5.425 6.814-8H9.195"/></g>`),
		g.Group(children),
	)
}