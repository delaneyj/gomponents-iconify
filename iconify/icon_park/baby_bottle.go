package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BabyBottle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M36 20H12V44H36V20Z"/><path d="M26 36H36"/><path d="M26 28H36"/><path d="M8 20H40"/><path d="M12 14H20.4V7.6C20.4 6.39815 21.6 4 24 4C26.4 4 27.6 6.39815 27.6 7.6V14H36"/></g>`),
		g.Group(children),
	)
}