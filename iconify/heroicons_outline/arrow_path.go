package heroicons_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowPath(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.5 12c0-1.232.046-2.453.138-3.662a4.006 4.006 0 0 1 3.7-3.7a48.678 48.678 0 0 1 7.324 0a4.006 4.006 0 0 1 3.7 3.7c.017.22.032.441.046.662M4.5 12l-3-3m3 3l3-3m12 3c0 1.232-.046 2.453-.138 3.662a4.006 4.006 0 0 1-3.7 3.7a48.657 48.657 0 0 1-7.324 0a4.006 4.006 0 0 1-3.7-3.7c-.017-.22-.032-.441-.046-.662M19.5 12l-3 3m3-3l3 3"/>`),
		g.Group(children),
	)
}