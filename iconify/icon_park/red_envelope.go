package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedEnvelope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M39 4H9V10L24 14L39 10V4Z"/><path d="M39 17V44H9V17"/><path d="M19 19L24 25L29 19"/><path d="M18 31H30"/><path d="M18 25H30"/><path d="M24 25V37"/></g>`),
		g.Group(children),
	)
}