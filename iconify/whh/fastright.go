package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fastright(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M589.49 467q19 18 19 44.5t-19 45.5l-451 448q-19 19-46 19t-46-19l-27-27q-19-18-19-44.5t19-45.5l379-376l-379-376q-19-19-19-45.5t19-45.5l27-26q19-19 46-19t46 19zm-154 511q-19-18-19-44.5t19-45.5l379-376l-379-376q-19-19-19-45.5t19-45.5l27-26q19-19 45.5-19t45.5 19l452 448q19 18 19 44.5t-19 45.5l-452 448q-19 19-45.5 19t-45.5-19z"/>`),
		g.Group(children),
	)
}