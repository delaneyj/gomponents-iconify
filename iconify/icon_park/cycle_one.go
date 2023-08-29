package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CycleOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 20C6 12 10 8 18 8"/><path d="M40 30C40 38 36 42 28 42"/><path fill="#2F88FF" d="M28 18C28 12.4772 32.4772 8 38 8H42V22H28V18Z"/><path fill="#2F88FF" d="M6 28H20V32C20 37.5228 15.5228 42 10 42H6V28Z"/></g>`),
		g.Group(children),
	)
}