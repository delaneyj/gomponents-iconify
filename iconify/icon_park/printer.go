package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Printer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M37 32H11V44H37V32Z"/><path stroke-linecap="round" d="M4 20H44V38H37.0173V32H10.9805V38H4V20Z" clip-rule="evenodd"/><path fill="#2F88FF" d="M38 4H10V20H38V4Z"/></g>`),
		g.Group(children),
	)
}