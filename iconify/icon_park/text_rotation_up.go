package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextRotationUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M29 35L21 31.3333M29 13L21 16.6667M21 16.6667L17 18.5L5 24L17 29.5L21 31.3333M21 16.6667V31.3333"/><path d="M37 42V6L43 12"/></g>`),
		g.Group(children),
	)
}