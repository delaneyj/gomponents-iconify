package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Loop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feLoop0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feLoop1" fill="currentColor"><path id="feLoop2" d="M6.876 15.124A6.002 6.002 0 0 0 17.658 14h2.09a8.003 8.003 0 0 1-14.316 2.568L3 19v-6h6l-2.124 2.124Zm10.249-6.249A6.004 6.004 0 0 0 6.34 10H4.25a8.005 8.005 0 0 1 14.32-2.57L21 5v6h-6l2.125-2.125Z"/></g></g>`),
		g.Group(children),
	)
}