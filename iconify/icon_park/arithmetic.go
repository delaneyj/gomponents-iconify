package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Arithmetic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M28 32.5H42"/><path d="M28 41.5H42"/><path d="M6 13H22"/><path d="M14 5L14 21"/><path d="M42 5L6 41"/></g>`),
		g.Group(children),
	)
}