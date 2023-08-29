package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AltArrowRightBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="m12.404 8.303l3.431 3.327c.22.213.22.527 0 .74l-6.63 6.43C8.79 19.201 8 18.958 8 18.43v-5.723l4.404-4.404Z"/><path d="M8 11.293V5.57c0-.528.79-.771 1.205-.37l2.481 2.406L8 11.293Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}