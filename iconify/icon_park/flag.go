package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M8 44H12H16"/><path d="M12 44V4"/><path fill="#2F88FF" d="M40 6H12V22H40L36 14L40 6Z"/></g>`),
		g.Group(children),
	)
}