package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tablet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.24 0H1.952A2.117 2.117 0 0 0 0 2.257V2.25v19.499a2.115 2.115 0 0 0 1.942 2.25h14.297a2.116 2.116 0 0 0 1.949-2.257v.007V2.25A2.122 2.122 0 0 0 16.246 0l-.008-.001zM9.096 22.5a1.498 1.498 0 1 1 .002-2.996a1.498 1.498 0 0 1-.002 2.996h-.004zm7.781-5.062a.564.564 0 0 1-.56.56H1.875a.564.564 0 0 1-.56-.56V2.813a.564.564 0 0 1 .56-.56h14.438a.564.564 0 0 1 .56.56z"/>`),
		g.Group(children),
	)
}