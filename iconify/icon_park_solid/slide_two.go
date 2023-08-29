package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlideTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSlideTwo0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path stroke="#fff" d="M4 8h40"/><path fill="#fff" fill-rule="evenodd" stroke="#fff" d="M8 8h32v26H8V8Z" clip-rule="evenodd"/><path stroke="#000" d="m31 18l3 3l-3 3m-14 0l-3-3l3-3"/><path stroke="#fff" d="m16 42l8-8l8 8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSlideTwo0)"/>`),
		g.Group(children),
	)
}