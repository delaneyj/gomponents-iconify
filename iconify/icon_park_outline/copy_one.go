package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CopyOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M13 38h28V16H30V4H13v34ZM30 4l11 12"/><path stroke-linejoin="round" d="M7 20v24h21"/><path d="M19 20h4m-4 8h12"/></g>`),
		g.Group(children),
	)
}