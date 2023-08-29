package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 12V4H40V12"/><path d="M40 36V44H24V36"/><path d="M24 24L4 24"/><path d="M32 34V14"/><path d="M12 16L4 24L12 32"/></g>`),
		g.Group(children),
	)
}