package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTxtOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M10 38V44H38V38"/><path stroke-linejoin="round" d="M38 20V14L30 4H10V20"/><path stroke-linejoin="round" d="M28 4V14H38"/><path d="M16 12H20"/><rect width="40" height="18" x="4" y="20" stroke-linejoin="round" rx="2"/><path stroke-linejoin="round" d="M21 25L27 33"/><path stroke-linejoin="round" d="M27 25L21 33"/><path stroke-linejoin="round" d="M13 25V33"/><path stroke-linejoin="round" d="M10 25H13H16"/><path stroke-linejoin="round" d="M35 25V33"/><path stroke-linejoin="round" d="M32 25H35H38"/></g>`),
		g.Group(children),
	)
}