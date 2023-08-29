package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.5 14.25a3.24 3.24 0 0 0-2.62 1.35L9.59 13a3.39 3.39 0 0 0 .16-1a3.39 3.39 0 0 0-.16-1l5.29-2.6a3.23 3.23 0 1 0-.63-1.9a2.94 2.94 0 0 0 .05.5L8.83 9.75a3.19 3.19 0 0 0-2.33-1a3.25 3.25 0 0 0 0 6.5a3.19 3.19 0 0 0 2.33-1L14.3 17a2.94 2.94 0 0 0-.05.51a3.25 3.25 0 1 0 3.25-3.25Zm0-9.5a1.75 1.75 0 1 1-1.75 1.75a1.76 1.76 0 0 1 1.75-1.75Zm-11 9A1.75 1.75 0 1 1 8.25 12a1.76 1.76 0 0 1-1.75 1.75Zm11 5.5a1.75 1.75 0 1 1 1.75-1.75a1.76 1.76 0 0 1-1.75 1.75Z"/>`),
		g.Group(children),
	)
}