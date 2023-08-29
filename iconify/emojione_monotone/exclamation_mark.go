package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExclamationMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M36.989 42.439H27.01L23 2h18z"/><ellipse cx="31.999" cy="54.354" fill="currentColor" rx="7.663" ry="7.646"/>`),
		g.Group(children),
	)
}