package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandNytimes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11.036 5.058a8 8 0 1 0 8.706 9.965"/><path d="M12 21V10l-7.5 4m13-11a2.5 2.5 0 1 1 0 5l-11-5a2.5 2.5 0 0 0-.67 4.91M9 12v8m7-7h-.01"/></g>`),
		g.Group(children),
	)
}