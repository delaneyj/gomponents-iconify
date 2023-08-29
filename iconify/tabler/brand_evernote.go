package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandEvernote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 8h5V3"/><path d="M17.9 19c.6-2.5 1.1-5.471 1.1-9c0-4.5-2-5-3-5c-1.906 0-3-.5-3.5-1c-.354-.354-.5-1-1.5-1H9L4 8c0 6 2.5 8 5 8c1 0 1.5-.5 2-1.5s1.414-.326 2.5 0c1.044.313 2.01.255 2.5.5c1 .5 2 1.5 2 3c0 .5 0 3-3 3s-3-3-1-3m1-8h1"/></g>`),
		g.Group(children),
	)
}