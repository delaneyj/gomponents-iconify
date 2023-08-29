package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StorageCardTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><rect width="32" height="40" x="8" y="4" stroke="#000" rx="2"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M14 4V14H34V4"/><rect width="20" height="20" x="14" y="24" fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M14 32H34"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M34 29L34 35"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M14 29L14 35"/></g>`),
		g.Group(children),
	)
}