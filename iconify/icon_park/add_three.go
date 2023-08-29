package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 32L24 16"/><path d="M42 27L42 21"/><path d="M6 27L6 21"/><path d="M14 6H8C6.89543 6 6 6.89543 6 8V14"/><path d="M34 6H40C41.1046 6 42 6.89543 42 8V14"/><path d="M34 42H40C41.1046 42 42 41.1046 42 40V34"/><path d="M14 42H8C6.89543 42 6 41.1046 6 40V34"/><path d="M27 6H21"/><path d="M32 24L16 24"/><path d="M27 42H21"/></g>`),
		g.Group(children),
	)
}