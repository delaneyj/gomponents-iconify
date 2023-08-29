package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeveloperGuideOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5ZM5 5v14h14V5h-2v6.125q0 .3-.25.438t-.5-.013l-1.225-.75q-.25-.15-.525-.15t-.525.15l-1.225.75q-.25.15-.5.013t-.25-.438V5H5Zm0 14V5v14Z"/>`),
		g.Group(children),
	)
}