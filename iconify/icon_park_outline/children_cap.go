package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChildrenCap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 32c0-15 5-24 8-24s6 6 6 6h8s3-6 6-6s8 9 8 24"/><rect width="40" height="8" x="4" y="32" rx="2"/></g>`),
		g.Group(children),
	)
}