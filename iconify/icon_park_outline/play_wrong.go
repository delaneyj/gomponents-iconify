package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayWrong(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M24 44C12.954 44 4 35.046 4 24S12.954 4 24 4s20 8.954 20 20m-11 9l8 8m0-8l-8 8"/><path d="M20 24v-6.928l6 3.464L32 24l-6 3.464l-6 3.464V24Z"/></g>`),
		g.Group(children),
	)
}