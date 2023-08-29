package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckboxBlank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M3 4h1V3h14v1h1v14h-1v1H4v-1H3V4m2 13h12V5H5v12Z"/>`),
		g.Group(children),
	)
}