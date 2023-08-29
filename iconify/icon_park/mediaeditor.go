package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mediaeditor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M35 32.1333V34C35 35.6569 33.6569 37 32 37H7C5.34315 37 4 35.6569 4 34V31.1333L44 12"/><path d="M35 15.8667V14C35 12.3431 33.6569 11 32 11H7C5.34315 11 4 12.3431 4 14V16.8667L44 36"/></g>`),
		g.Group(children),
	)
}