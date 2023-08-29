package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleDoubleLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8.207 12l3.146-3.146a.5.5 0 0 0-.707-.707l-3.5 3.5a.5.5 0 0 0 0 .707l3.5 3.5a.498.498 0 0 0 .707 0a.5.5 0 0 0 0-.708L8.208 12zm9.147 3.146L14.207 12l3.146-3.146a.5.5 0 0 0-.707-.707l-3.5 3.5a.5.5 0 0 0 0 .707l3.5 3.5a.498.498 0 0 0 .707 0a.5.5 0 0 0 0-.708z"/>`),
		g.Group(children),
	)
}