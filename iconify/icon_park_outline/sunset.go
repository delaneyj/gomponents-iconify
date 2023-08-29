package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sunset(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linejoin="round" d="m19 8l14 26H5L19 8Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m29 26l5-6l9 14H32m-22 7h28"/><circle cx="38" cy="10" r="3"/></g>`),
		g.Group(children),
	)
}