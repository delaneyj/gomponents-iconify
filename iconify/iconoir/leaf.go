package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Leaf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 21s.5-4.5 4-8.5"/><path d="m19.13 4.242l.594 6.175c.374 3.886-2.54 7.346-6.425 7.72c-3.813.367-7.267-2.42-7.634-6.233a6.936 6.936 0 0 1 6.239-7.569l6.571-.632a.6.6 0 0 1 .655.54Z"/></g>`),
		g.Group(children),
	)
}