package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeinConnect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.991 19.843a3.616 3.616 0 1 1 .564-1.94h-4.13M9.5 17.904a3.617 3.617 0 1 0 7.233 0h0a3.617 3.617 0 1 0-7.233 0v-6.318m19.437 9.83v-5.148c0-1.725 2.246-2.563 3.328-1.037c1.128 1.59 3.018 5.428 3.018 5.428c.847 1.405 3.217.5 3.217-.992v-5.276m-11.355 0v7.025m1.594 9.527l-9.478 5.472V25.471l9.478 5.472z"/><rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/>`),
		g.Group(children),
	)
}