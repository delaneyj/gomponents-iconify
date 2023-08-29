package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 23v-4H7q-.825 0-1.413-.588T5 17V7H1V5h4V1h2v16h16v2h-4v4h-2Zm0-8V7H9V5h8q.825 0 1.413.588T19 7v8h-2Z"/>`),
		g.Group(children),
	)
}