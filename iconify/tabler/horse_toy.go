package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HorseToy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3.5 17.5c5.667 4.667 11.333 4.667 17 0"/><path d="M19 18.5L17 10l1-2l2 1l1.5-1.5L19 3c-5.052.218-5.99 3.133-7 6H6a3 3 0 0 0-3 3m2 6.5L7 9"/><path d="m8 20l2-5h4l2 5"/></g>`),
		g.Group(children),
	)
}