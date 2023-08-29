package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailUnpacking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path d="M44 18V39.8182C44 41.0232 43.1046 42 42 42H6C4.89543 42 4 41.0232 4 39.8182V18L24 34L44 18Z"/><path stroke-linecap="round" d="M4 17.7839L24 4L44 17.7839"/><path fill="#2F88FF" d="M34 18H14V26L24 34L34 26V18Z"/></g>`),
		g.Group(children),
	)
}