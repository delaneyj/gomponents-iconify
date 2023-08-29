package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sporting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><circle cx="24" cy="8" r="4"/><path stroke-linecap="round" stroke-linejoin="round" d="M7 18h12v16m22-16H29v26"/></g>`),
		g.Group(children),
	)
}