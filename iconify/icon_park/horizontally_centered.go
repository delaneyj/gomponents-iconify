package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HorizontallyCentered(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><rect width="16" height="16" x="16" y="16" fill="#2F88FF" stroke-linejoin="round"/><path d="M5 40L5 8"/><path d="M43 40L43 8"/></g>`),
		g.Group(children),
	)
}