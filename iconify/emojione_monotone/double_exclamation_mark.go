package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleExclamationMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M22.989 42.439H13.01L9 2h18z"/><ellipse cx="17.999" cy="54.354" fill="currentColor" rx="7.663" ry="7.646"/><path fill="currentColor" d="M50.989 42.439H41.01L37 2h18z"/><ellipse cx="45.999" cy="54.354" fill="currentColor" rx="7.663" ry="7.646"/>`),
		g.Group(children),
	)
}