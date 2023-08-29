package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Edit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M7 42H43"/><path fill="#2F88FF" d="M11 26.7199V34H18.3172L39 13.3081L31.6951 6L11 26.7199Z"/></g>`),
		g.Group(children),
	)
}