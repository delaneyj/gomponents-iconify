package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hammer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.634 11.056L2.148 19.54l2.122 2.121l8.485-8.485"/><path d="m10.634 11.056l1.414-1.415s.354-3.182-3.182-6.717l1.06-1.06l8.486 5.656l-1.06 1.06l1.413 1.415l1.061-1.06l2.475 2.474l-4.95 4.95l-2.475-2.475l1.061-1.06l-1.414-1.415l-1.768 1.768l-2.121-2.121Z"/></g>`),
		g.Group(children),
	)
}