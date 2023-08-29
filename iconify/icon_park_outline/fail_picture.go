package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FailPicture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><rect width="34" height="34" x="7" y="7" rx="3"/><path stroke-linecap="round" d="m41 26l-6.612-4.96a2 2 0 0 0-2.614.187L27 26M7 28l11-10M6 6l36 36"/></g>`),
		g.Group(children),
	)
}