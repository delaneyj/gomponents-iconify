package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Buzzfeed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M24 12c0 6.627-5.373 12-12 12S0 18.627 0 12S5.373 0 12 0s12 5.373 12 12zm-4.148-.273l-.977-6.94l-6.5 2.624l2.575 1.487l-2.435 4.215L8.3 10.68l-4.153 7.19l2.327 1.346l2.812-4.868L13.5 16.78l3.777-6.54l2.575 1.487z"/>`),
		g.Group(children),
	)
}