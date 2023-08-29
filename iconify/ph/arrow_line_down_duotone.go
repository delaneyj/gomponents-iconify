package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLineDownDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m200 112l-72 72l-72-72Z" opacity=".2"/><path d="M122.34 189.66a8 8 0 0 0 11.32 0l72-72A8 8 0 0 0 200 104h-64V32a8 8 0 0 0-16 0v72H56a8 8 0 0 0-5.66 13.66ZM180.69 120L128 172.69L75.31 120ZM224 216a8 8 0 0 1-8 8H40a8 8 0 0 1 0-16h176a8 8 0 0 1 8 8Z"/></g>`),
		g.Group(children),
	)
}