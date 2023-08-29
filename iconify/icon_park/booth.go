package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Booth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M4 5H44V13L42.6015 13.8391C40.3847 15.1692 37.6153 15.1692 35.3985 13.8391L34 13L32.6015 13.8391C30.3847 15.1692 27.6153 15.1692 25.3985 13.8391L24 13L22.6015 13.8391C20.3847 15.1692 17.6153 15.1692 15.3985 13.8391L14 13L12.6015 13.8391C10.3847 15.1692 7.61531 15.1692 5.39853 13.8391L4 13V5Z"/><rect width="36" height="18" x="6" y="25"/><path d="M9 16V25"/><path d="M39 16V25"/></g>`),
		g.Group(children),
	)
}