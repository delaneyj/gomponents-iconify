package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 13.25c.5-6 5.5-7.5 8-7v-3.5L14.25 8l-4.5 5.25v-3.5c-2.5-.5-6.5.5-8 3.5z"/>`),
		g.Group(children),
	)
}