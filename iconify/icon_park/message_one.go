package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" stroke-linejoin="round" d="M4 6H44V36H29L24 41L19 36H4V6Z"/><path stroke="#fff" d="M23 21H25.0025"/><path stroke="#fff" d="M33.001 21H34.9999"/><path stroke="#fff" d="M13.001 21H14.9999"/></g>`),
		g.Group(children),
	)
}