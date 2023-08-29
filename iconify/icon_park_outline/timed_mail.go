package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimedMail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M44 35V9H4v28h22"/><circle cx="35" cy="35" r="9"/><path stroke-linecap="round" stroke-linejoin="round" d="M34 32v4h4M4 9l20 13L44 9"/></g>`),
		g.Group(children),
	)
}