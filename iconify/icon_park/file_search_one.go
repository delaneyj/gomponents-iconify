package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileSearchOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M40 23V14L31 4H10C8.89543 4 8 4.89543 8 6V42C8 43.1046 8.89543 44 10 44H22"/><circle cx="31" cy="34" r="6" fill="#2F88FF"/><path stroke-linecap="round" stroke-linejoin="round" d="M36 38L41 42"/><path stroke-linecap="round" stroke-linejoin="round" d="M30 4V14H40"/></g>`),
		g.Group(children),
	)
}