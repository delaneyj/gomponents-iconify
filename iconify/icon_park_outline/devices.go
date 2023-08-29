package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Devices(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M23 43h20V5H14v10"/><path d="M5 15h18v28H5V15Z"/><path stroke-linecap="round" d="M13 37h2m13 0h2"/></g>`),
		g.Group(children),
	)
}