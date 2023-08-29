package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" fill-rule="evenodd" d="M11 14L25 4V44H11V14Z" clip-rule="evenodd"/><path d="M25 13L39 23V44"/><path d="M4 44H44"/></g>`),
		g.Group(children),
	)
}