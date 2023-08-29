package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WhiteExclamationMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<g fill="#a6aeb0"><path d="M37 42.4H27L23 2h18z"/><ellipse cx="32" cy="54.4" rx="7.7" ry="7.6"/></g>`),
		g.Group(children),
	)
}