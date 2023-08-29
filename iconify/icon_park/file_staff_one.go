package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileStaffOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M40 23V14L31 4H10C8.89543 4 8 4.89543 8 6V42C8 43.1046 8.89543 44 10 44H20"/><circle cx="34" cy="32" r="4" fill="#2F88FF"/><path d="M42 44C42 39.5817 38.4183 36 34 36C29.5817 36 26 39.5817 26 44"/><path d="M30 4V14H40"/></g>`),
		g.Group(children),
	)
}