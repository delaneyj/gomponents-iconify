package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapDraw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M17 12L4 6v30l13 6l14-6l13 6V12L31 6l-14 6Zm14-6v30M17 12v30"/><path d="m10.5 9l6.5 3l14-6l6.5 3m-27 30l6.5 3l14-6l6.5 3"/></g>`),
		g.Group(children),
	)
}