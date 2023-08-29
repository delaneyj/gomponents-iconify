package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shaving(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="20" height="27" x="14" y="17" rx="2"/><path d="M19 12h10v5H19zm0 0V9c0-1 0-5 8-5h9v5h-7v3"/></g>`),
		g.Group(children),
	)
}