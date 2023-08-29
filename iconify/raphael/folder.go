package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Folder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28.625 26.75h-26.5V8.375H3.25c1.75 0 .747-3.125 3-3.125h5.125c2.25 0 1.25 3.125 3 3.125h14.25V26.75z"/>`),
		g.Group(children),
	)
}