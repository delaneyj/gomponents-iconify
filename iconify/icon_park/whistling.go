package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Whistling(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4"><path d="M27 11L4 11V19H15C14.35 20.55 14 22.21 14 24C14 31.21 19.79 37 27 37C34.2 37 40 31.21 40 24C40 16.79 34.21 11 27 11Z"/><path d="M27 11V17"/><path d="M40 24H44"/></g>`),
		g.Group(children),
	)
}