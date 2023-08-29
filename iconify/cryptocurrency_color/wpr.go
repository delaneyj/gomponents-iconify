package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wpr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#ffe600"/><path fill="#fff" d="M3.993 5.425a16.096 16.096 0 0 1 2.65-2.405l4.452 15.918l1.973-8.438h3.74l1.972 8.438l2.337-8.438h3.749l.035.15l4.074 14.714a16.092 16.092 0 0 1-2.55 2.775l-3.437-12.41l-1.96 7.078l-.035.15h-4.186l-1.869-6.932l-1.87 6.932H8.882l-.035-.15z"/></g>`),
		g.Group(children),
	)
}