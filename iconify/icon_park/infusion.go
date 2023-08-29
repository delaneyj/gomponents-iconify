package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Infusion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M38 30C38 37.732 31.732 44 24 44C16.268 44 10 37.732 10 30C10 22.268 24 4 24 4C24 4 38 22.268 38 30Z"/><path stroke="#fff" stroke-linecap="round" d="M18 30H30"/><path stroke="#fff" stroke-linecap="round" d="M24 24V36"/></g>`),
		g.Group(children),
	)
}