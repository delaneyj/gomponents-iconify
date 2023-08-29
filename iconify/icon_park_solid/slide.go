package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Slide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSlide0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M41 5.5H7v28h34v-28Z"/><path stroke="#fff" stroke-linecap="round" d="m16 41.5l8-8l8 8"/><path stroke="#000" stroke-linecap="round" d="m13.924 24.663l5.642-5.508l4.442 4.345l9.959-9.98"/><path stroke="#fff" stroke-linecap="round" d="M4 5.5h40"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSlide0)"/>`),
		g.Group(children),
	)
}