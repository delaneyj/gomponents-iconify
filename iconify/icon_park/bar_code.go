package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 20H44"/><path d="M6 6V14"/><path d="M6 25.9956V37.9992"/><path d="M20.4 6V14"/><path d="M20.4 26V42"/><path d="M34.8 6V14"/><path d="M42 6V14"/><path d="M34.8 26V34"/><path d="M13.2 6V14"/><path d="M13.2 26V34"/><path d="M27.6 6V14"/><path d="M27.6 26V34"/><path d="M42 26V38"/></g>`),
		g.Group(children),
	)
}