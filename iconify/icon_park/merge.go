package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Merge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="20" height="22" x="4" y="24" rx="2" transform="rotate(-45 4 24)"/><rect width="20" height="20" x="16" y="24" rx="2" transform="rotate(-45 16 24)"/></g>`),
		g.Group(children),
	)
}