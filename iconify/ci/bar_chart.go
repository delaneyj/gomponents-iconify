package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 21h-3V11h3v10Zm-5 0h-3V8h3v13Zm-5 0H8V5h3v16Zm-5 0H3v-8h3v8Z"/>`),
		g.Group(children),
	)
}