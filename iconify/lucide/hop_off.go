package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HopOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17.5 5.5C19 7 20.5 9 21 11c-1.323.265-2.646.39-4.118.226M5.5 17.5C7 19 9 20.5 11 21c.5-2.5.5-5-1-8.5m7.5 5c-2.5 0-4 0-6-1m8.5-5c1 1.5 2 3.5 2 4.5m-10.5 4c1.5 1 3.5 2 4.5 2c.5-1.5 0-3-.5-4.5M22 22c-2 0-3.5-.5-5.5-1.5"/><path d="M4.783 4.782C1.073 8.492 1 14.5 5 18c1-1 2-4.5 1.5-6.5c1.5 1 4 1 5.5.5M8.227 2.57C11.578 1.335 15.453 2.089 18 5c-.88.88-3.7 1.761-5.726 1.618M2 2l20 20"/></g>`),
		g.Group(children),
	)
}