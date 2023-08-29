package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Exposure(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M277 323h-42v-43h42v-43h43v43h43v43h-43v42h-43v-42zM384 3q18 0 30.5 12.5T427 45v342q0 17-12.5 29.5T384 429H43q-18 0-30.5-12.5T0 387V45q0-17 12.5-29.5T43 3h341zM64 67v42h128V67H64zm320 320V45L43 387h341z"/>`),
		g.Group(children),
	)
}