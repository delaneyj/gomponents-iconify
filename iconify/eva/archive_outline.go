package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaArchiveOutline0"><g id="evaArchiveOutline1"><g id="evaArchiveOutline2" fill="currentColor"><path d="M21 6a3 3 0 0 0-3-3H6a3 3 0 0 0-2 5.22V18a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V8.22A3 3 0 0 0 21 6ZM6 5h12a1 1 0 0 1 0 2H6a1 1 0 0 1 0-2Zm12 13a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V9h12Z"/><rect width="6" height="2" x="9" y="12" rx=".87" ry=".87"/></g></g></g>`),
		g.Group(children),
	)
}