package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitBranch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path fill="#2F88FF" d="M44 44V4H24V17L33 26V44H44Z"/><path fill="#2F88FF" d="M4 4V44H25V30L16 21V4H4Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M44 44V4H24V17L33 26V44H44Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 4V44H25V30L16 21V4H4Z"/></g>`),
		g.Group(children),
	)
}