package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M6.5 12.5a1.5 1.5 0 0 1 0-3h10a1.5 1.5 0 0 1 0 3h-10Z" opacity=".2"/><path d="M5 10.5a.5.5 0 0 1 0-1h10a.5.5 0 0 1 0 1H5Z"/></g>`),
		g.Group(children),
	)
}