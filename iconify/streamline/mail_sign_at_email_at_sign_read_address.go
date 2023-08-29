package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailSignAtEmailAtSignReadAddress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10.05 7A3 3 0 1 1 7 4a3 3 0 0 1 3.05 3Z"/><path d="M10.05 7v1.3c0 3.49 5.47.2 2.6-4.54A6.59 6.59 0 0 0 7 .5A6.5 6.5 0 1 0 9.52 13"/></g>`),
		g.Group(children),
	)
}