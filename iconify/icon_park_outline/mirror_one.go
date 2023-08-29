package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MirrorOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><circle cx="24" cy="20" r="16"/><path stroke-linecap="round" stroke-linejoin="round" d="M24 36v8m-10 0h20"/></g>`),
		g.Group(children),
	)
}