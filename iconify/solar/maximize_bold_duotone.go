package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MaximizeBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M8.607 22a.75.75 0 0 1-.75.75H2a.75.75 0 0 1-.75-.75v-5.857a.75.75 0 1 1 1.5 0v4.046l5.72-5.72a.75.75 0 0 1 1.06 1.061l-5.72 5.72h4.047a.75.75 0 0 1 .75.75Z" opacity=".6"/><path d="M15.393 2a.75.75 0 0 1 .75-.75H22a.75.75 0 0 1 .75.75v5.857a.75.75 0 0 1-1.5 0V3.811l-5.72 5.72a.75.75 0 1 1-1.06-1.061l5.72-5.72h-4.047a.75.75 0 0 1-.75-.75Z"/></g>`),
		g.Group(children),
	)
}