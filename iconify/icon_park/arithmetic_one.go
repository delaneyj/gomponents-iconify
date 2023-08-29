package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArithmeticOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M28 31.5H42"/><path d="M28 39.5H42"/><path d="M7.34281 40.6568L18.6565 29.3431"/><path d="M7.3428 29.3433L18.6565 40.657"/><path d="M28 12.5H42"/><path d="M6 12.5H20"/><path d="M13 5.5V19.5"/></g>`),
		g.Group(children),
	)
}