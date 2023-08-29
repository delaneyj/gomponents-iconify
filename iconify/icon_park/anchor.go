package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Anchor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><line x1="14" x2="34" y1="8" y2="8" stroke-linecap="round"/><line x1="14" x2="34" y1="8" y2="8" stroke-linecap="round"/><line x1="14" x2="34" y1="40" y2="40" stroke-linecap="round"/><rect width="8" height="8" x="36" y="4" fill="#2F88FF" stroke-linejoin="round" rx="2"/><rect width="8" height="8" x="4" y="4" fill="#2F88FF" stroke-linejoin="round" rx="2"/><rect width="8" height="8" x="36" y="36" fill="#2F88FF" stroke-linejoin="round" rx="2"/><rect width="8" height="8" x="4" y="36" fill="#2F88FF" stroke-linejoin="round" rx="2"/><line x1="40" x2="40" y1="14" y2="34" stroke-linecap="round"/><line x1="8" x2="8" y1="14" y2="34" stroke-linecap="round"/></g>`),
		g.Group(children),
	)
}