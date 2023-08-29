package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pyramid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path d="M21 12L38 42H4L21 12Z" clip-rule="evenodd"/><path stroke-linecap="round" d="M36.5 42H44L36 28L33 33"/><path stroke-linecap="round" d="M21 12L13 42"/></g>`),
		g.Group(children),
	)
}