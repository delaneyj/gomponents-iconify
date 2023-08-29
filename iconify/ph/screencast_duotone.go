package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreencastDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M224 56v144a8 8 0 0 1-8 8H48a16 16 0 0 0-16-16V56a8 8 0 0 1 8-8h176a8 8 0 0 1 8 8Z" opacity=".2"/><path d="M232 56v144a16 16 0 0 1-16 16h-72a8 8 0 0 1 0-16h72V56H40v40a8 8 0 0 1-16 0V56a16 16 0 0 1 16-16h176a16 16 0 0 1 16 16ZM32 184a8 8 0 0 0 0 16a8 8 0 0 1 8 8a8 8 0 0 0 16 0a24 24 0 0 0-24-24Zm0-32a8 8 0 0 0 0 16a40 40 0 0 1 40 40a8 8 0 0 0 16 0a56.06 56.06 0 0 0-56-56Zm0-32a8 8 0 0 0 0 16a72.08 72.08 0 0 1 72 72a8 8 0 0 0 16 0a88.1 88.1 0 0 0-88-88Z"/></g>`),
		g.Group(children),
	)
}