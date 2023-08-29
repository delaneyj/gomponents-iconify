package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IncomeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M35 38L30 33L34.9996 28"/><path d="M43 38L38 33L42.9996 28"/><path d="M43 22V9C43 7.89543 42.1046 7 41 7H7C5.89543 7 5 7.89543 5 9V39C5 40.1046 5.89543 41 7 41H28.4706"/><path d="M13 15L18 21L23 15"/><path d="M12 27H24"/><path d="M12 21H24"/><path d="M18 21V33"/></g>`),
		g.Group(children),
	)
}