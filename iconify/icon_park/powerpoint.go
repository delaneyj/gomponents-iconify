package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Powerpoint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#2F88FF" stroke="#000" rx="3"/><rect width="10" height="10" x="19" y="15" stroke="#fff" stroke-linecap="round" stroke-linejoin="round"/><path stroke="#fff" stroke-linecap="round" d="M19 33L19 15"/></g>`),
		g.Group(children),
	)
}