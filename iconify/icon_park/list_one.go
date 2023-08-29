package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M18 6H34V12H18V6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M18 21H38V27H18V21Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M18 36H44V42H18V36Z"/><circle cx="8" cy="9" r="2"/><circle cx="8" cy="24" r="2"/><circle cx="8" cy="39" r="2"/></g>`),
		g.Group(children),
	)
}