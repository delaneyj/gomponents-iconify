package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextAlignLeftDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M216 64v112a8 8 0 0 1-8 8H40V64Z" opacity=".2"/><path d="M32 64a8 8 0 0 1 8-8h176a8 8 0 0 1 0 16H40a8 8 0 0 1-8-8Zm8 48h128a8 8 0 0 0 0-16H40a8 8 0 0 0 0 16Zm176 24H40a8 8 0 0 0 0 16h176a8 8 0 0 0 0-16Zm-48 40H40a8 8 0 0 0 0 16h128a8 8 0 0 0 0-16Z"/></g>`),
		g.Group(children),
	)
}