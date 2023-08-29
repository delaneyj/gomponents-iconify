package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M34 28L44 24"/><path fill="#2F88FF" d="M4 28H34L33.5613 31.8024C33.1537 35.3345 30.163 38 26.6074 38H11.3926C7.83703 38 4.84629 35.3345 4.43873 31.8024L4 28Z"/><path d="M19 10V20"/><path d="M11 12V18"/><path d="M27 12V18"/></g>`),
		g.Group(children),
	)
}