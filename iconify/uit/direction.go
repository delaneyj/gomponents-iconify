package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Direction(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.854 9.854L12 6.707l3.146 3.147a.498.498 0 0 0 .707 0a.5.5 0 0 0 0-.707l-3.5-3.5a.5.5 0 0 0-.706 0l-3.5 3.5a.5.5 0 0 0 .707.707zm6.292 4.292L12 17.293l-3.147-3.146a.5.5 0 0 0-.707.707l3.5 3.5a.498.498 0 0 0 .707 0l3.5-3.5a.5.5 0 0 0-.707-.708z"/>`),
		g.Group(children),
	)
}