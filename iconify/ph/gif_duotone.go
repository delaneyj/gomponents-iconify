package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GifDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M232 56v144a16 16 0 0 1-16 16H40a16 16 0 0 1-16-16V56a16 16 0 0 1 16-16h176a16 16 0 0 1 16 16Z" opacity=".2"/><path d="M208 88a8 8 0 0 1-8 8h-24v24h16a8 8 0 0 1 0 16h-16v32a8 8 0 0 1-16 0V88a8 8 0 0 1 8-8h32a8 8 0 0 1 8 8Zm-72-8a8 8 0 0 0-8 8v80a8 8 0 0 0 16 0V88a8 8 0 0 0-8-8Zm-32 40H88a8 8 0 0 0 0 16h8v8a16 16 0 0 1-32 0v-32a16 16 0 0 1 27.93-10.67a8 8 0 1 0 11.92-10.66A32 32 0 0 0 48 112v32a32 32 0 0 0 64 0v-16a8 8 0 0 0-8-8Z"/></g>`),
		g.Group(children),
	)
}