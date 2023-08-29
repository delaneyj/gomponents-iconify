package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TShirt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M9 9L18 4H30L39 9L43 24L35 30V44H13V30L5 24L9 9Z"/><path d="M13 31L13 24"/><path d="M35 31L35 24"/></g>`),
		g.Group(children),
	)
}