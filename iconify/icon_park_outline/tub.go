package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tub(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path d="M40 23V12a7 7 0 0 0-7-7v0a7 7 0 0 0-7 7v1m14 16v-6H8v6a8 8 0 0 0 8 8h16a8 8 0 0 0 8-8Zm3-6H5"/><path stroke-linejoin="round" d="m17 37l-4 6m18-6l4 6"/></g>`),
		g.Group(children),
	)
}