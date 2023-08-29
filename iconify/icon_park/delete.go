package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Delete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M9 10V44H39V10H9Z"/><path stroke="#fff" stroke-linecap="round" d="M20 20V33"/><path stroke="#fff" stroke-linecap="round" d="M28 20V33"/><path stroke="#000" stroke-linecap="round" d="M4 10H44"/><path fill="#2F88FF" stroke="#000" d="M16 10L19.289 4H28.7771L32 10H16Z"/></g>`),
		g.Group(children),
	)
}