package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cardmembership(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23 3v18H1V3h22Zm-2 2H3v4h18V5Zm0 6h-2v6.766l-3.5-2.1l-3.5 2.1V11H3v8h18v-8Zm-4 0h-3v3.234l1.5-.9l1.5.9V11Z"/>`),
		g.Group(children),
	)
}