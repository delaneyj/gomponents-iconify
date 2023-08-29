package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sweater(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M14 37H6V24L11 9L19 4H24H29L38 9L42 24V37H34V44H24H14V37Z"/><path d="M34 28V37"/><path d="M14 28V37"/></g>`),
		g.Group(children),
	)
}