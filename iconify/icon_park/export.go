package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Export(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M42 27C42 33 38 43 24 43C10 43 6 33 6 27"/><path d="M24.0078 5.10059V33.0001"/><path d="M12 17L24 5L36 17"/></g>`),
		g.Group(children),
	)
}