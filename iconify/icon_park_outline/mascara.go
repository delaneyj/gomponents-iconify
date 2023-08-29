package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mascara(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M28 16h12v28H28zM8 28h12v16H8zm6-24v24"/><path d="M20 36H8"/><path stroke-linejoin="round" d="M20 32v8M8 32v8M18 9h-8m10 6H8m10 6h-8"/></g>`),
		g.Group(children),
	)
}