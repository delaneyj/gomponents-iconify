package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AltArrowUpBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="m8.303 11.596l3.327-3.431a.499.499 0 0 1 .74 0l6.43 6.63c.401.414.158 1.205-.37 1.205h-5.723l-4.404-4.404Z"/><path d="M11.293 16H5.57c-.528 0-.771-.791-.37-1.205l2.406-2.482L11.293 16Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}