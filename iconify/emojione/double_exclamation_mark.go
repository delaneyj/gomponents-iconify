package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleExclamationMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<g fill="#ff5a79"><path d="M23 42.4H13L9 2h18z"/><ellipse cx="18" cy="54.4" rx="7.7" ry="7.6"/><path d="M51 42.4H41L37 2h18z"/><ellipse cx="46" cy="54.4" rx="7.7" ry="7.6"/></g>`),
		g.Group(children),
	)
}