package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Textarea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M16 4C23 4 24 10 24 12C24 14 24 34 24 36C24 38 23 44 16 44"/><path d="M32 4C26 4 24 10 24 12C24 14 24 34 24 36C24 38 25 44 32 44"/><path d="M17 24L31 24"/></g>`),
		g.Group(children),
	)
}