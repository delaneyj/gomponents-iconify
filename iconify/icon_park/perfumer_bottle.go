package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PerfumerBottle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="38" height="24" x="5" y="17" fill="#2F88FF" stroke="#000" rx="2"/><rect width="20" height="10" x="14" y="7" fill="#2F88FF" stroke="#000"/><rect width="12" height="8" x="18" y="25" fill="#43CCF8" stroke="#fff"/><path stroke="#fff" d="M30 29H43"/><path stroke="#fff" d="M5 29H18"/><path stroke="#000" d="M5 24V34"/><path stroke="#000" d="M43 24V34"/></g>`),
		g.Group(children),
	)
}