package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuestionMarkSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" stroke-width="1.5" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-14Z"/><path stroke-width="2.25" d="M12 16h.01v.01H12z"/><path stroke-linecap="round" stroke-width="1.5" d="M10.586 7.586c.39-.39.9-.585 1.41-.586a1.991 1.991 0 0 1 1.418.586c.39.39.586.902.586 1.414a1.99 1.99 0 0 1-.586 1.414a1.993 1.993 0 0 1-1.418.586L12 12"/></g>`),
		g.Group(children),
	)
}