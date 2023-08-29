package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Commit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 15h-8.09a5.992 5.992 0 0 0-11.82 0H2v2h8.09a5.992 5.992 0 0 0 11.82 0H30Zm-14 5a4 4 0 1 1 4-4a4.005 4.005 0 0 1-4 4Z"/>`),
		g.Group(children),
	)
}