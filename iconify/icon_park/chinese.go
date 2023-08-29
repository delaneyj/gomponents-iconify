package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chinese(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#2F88FF" stroke="#000" stroke-linejoin="round" rx="3"/><rect width="20" height="10" x="14" y="18" fill="#43CCF8" stroke="#fff" stroke-linejoin="round"/><path stroke="#fff" d="M24 14V35"/></g>`),
		g.Group(children),
	)
}