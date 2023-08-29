package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PercentageFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M9 14a5 5 0 1 1 5-5a5.005 5.005 0 0 1-5 5ZM4 26.586L26.585 4L28 5.415L5.414 28zM23 28a5 5 0 1 1 5-5a5.005 5.005 0 0 1-5 5Z"/>`),
		g.Group(children),
	)
}