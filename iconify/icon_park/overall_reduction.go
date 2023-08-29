package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OverallReduction(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M15 15H33V33H15V15Z"/><path d="M11 43V37H5"/><path d="M37 43V37H43"/><path d="M11 5V11H5"/><path d="M37 5V11H43"/></g>`),
		g.Group(children),
	)
}