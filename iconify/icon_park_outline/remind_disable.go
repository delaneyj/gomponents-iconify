package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemindDisable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path d="M24 4c-7.732 0-14 6.268-14 14v20h28V18c0-7.732-6.268-14-14-14Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M10 38V18c0-7.732 6.268-14 14-14s14 6.268 14 14v20M4 38h40m-20 6a5 5 0 0 0 5-5v-1H19v1a5 5 0 0 0 5 5Zm-5-27l10 10m0-10L19 27"/></g>`),
		g.Group(children),
	)
}