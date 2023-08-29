package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Increase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M41 25c0 9.941-8.059 18-18 18S5 34.941 5 25S13.059 7 23 7"/><path d="M12 28.5c8.5 0 12-.5 19-9.5"/><path d="M23 19h8v8m0-22v4.5M43.5 17H39m1.89-10L37 10.89"/></g>`),
		g.Group(children),
	)
}