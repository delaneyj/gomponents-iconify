package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path stroke="#000" stroke-linecap="round" d="M8 11L40 11"/><path stroke="#000" stroke-linecap="round" d="M18 5L30 5"/><path fill="#2F88FF" stroke="#000" d="M12 17H36V40C36 41.6569 34.6569 43 33 43H15C13.3431 43 12 41.6569 12 40V17Z"/><path stroke="#fff" stroke-linecap="round" d="M20 25L28 33"/><path stroke="#fff" stroke-linecap="round" d="M28 25L20 33"/></g>`),
		g.Group(children),
	)
}