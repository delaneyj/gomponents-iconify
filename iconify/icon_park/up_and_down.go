package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpAndDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M12 36V44H44V12H36V20H28V28H20V36H12Z"/><path d="M18 13L27 4"/><path d="M21 4H27V10"/><path d="M10 27H4V21"/><path d="M13 18L4 27"/></g>`),
		g.Group(children),
	)
}