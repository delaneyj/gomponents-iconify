package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleDoubleRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.353 8.147a.5.5 0 0 0-.707.707L9.793 12l-3.147 3.146a.5.5 0 1 0 .707.708l3.5-3.5a.5.5 0 0 0 0-.707l-3.5-3.5zm9.5 3.5l-3.5-3.5a.5.5 0 0 0-.707.707L15.793 12l-3.147 3.146a.5.5 0 1 0 .707.708l3.5-3.5a.5.5 0 0 0 0-.707z"/>`),
		g.Group(children),
	)
}