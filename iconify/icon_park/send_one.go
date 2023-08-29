package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path d="M42 6L4 20.1383L24 24.0083L29.0052 44L42 6Z"/><path stroke-linecap="round" d="M24.0083 24.0084L29.6651 18.3516"/></g>`),
		g.Group(children),
	)
}