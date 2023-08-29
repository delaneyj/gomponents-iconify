package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartbeatAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.909 10.928H24v2.02h-2.091zm-3.635 0h2.091v2.02h-2.091zm-6.022.001l-.657 2.215l-3.772-8.343l-1.954 6.17H0v2.02h7.346l.81-2.551l3.834 8.486l1.76-5.978h2.973v-2.02z"/>`),
		g.Group(children),
	)
}