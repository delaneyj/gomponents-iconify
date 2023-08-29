package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForTriangularFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ccc" d="M26.2 61.5c0 1.4-4.3 2.5-9.7 2.5c-5.4 0-9.7-1.1-9.7-2.5s4.3-2.5 9.7-2.5c5.3 0 9.7 1.2 9.7 2.5"/><path fill="#ed4c5c" d="M56.1 17.7C54.2 16.8 23 2.1 18.4 0v39.5c4.5-2.1 35.7-16.8 37.6-17.7c1.6-.7 1.6-3.4.1-4.1"/><path fill="#d3976e" d="M14.5 0h3.9v61.5h-3.9z"/>`),
		g.Group(children),
	)
}