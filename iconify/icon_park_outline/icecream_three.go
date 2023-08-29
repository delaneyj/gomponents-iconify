package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IcecreamThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path d="M14 14c0-5.523 4.477-10 10-10s10 4.477 10 10v17.73a.27.27 0 0 1-.27.27H14.27a.27.27 0 0 1-.27-.27V14Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M21 32v9a3 3 0 0 0 3 3v0a3 3 0 0 0 3-3v-9M14 16v0a7.071 7.071 0 0 0 10 0v0a7.071 7.071 0 0 1 10 0v0"/></g>`),
		g.Group(children),
	)
}