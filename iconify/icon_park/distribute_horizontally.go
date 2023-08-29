package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DistributeHorizontally(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><rect width="28" height="12" x="30" y="10" fill="#2F88FF" transform="rotate(90 30 10)"/><path stroke-linecap="round" d="M40 6V42"/><path stroke-linecap="round" d="M8 6V42"/></g>`),
		g.Group(children),
	)
}