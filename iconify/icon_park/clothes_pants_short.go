package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClothesPantsShort(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M6 36L8.00001 12H40L42 36H26.8421L24 25L21.1579 36H6Z"/><path stroke="#fff" d="M24 12L27 19"/><path stroke="#fff" d="M24 12L20 19.5"/><path stroke="#000" d="M18 12H30"/></g>`),
		g.Group(children),
	)
}