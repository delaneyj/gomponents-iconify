package ant_design

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EllipsisOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M176 511a56 56 0 1 0 112 0a56 56 0 1 0-112 0zm280 0a56 56 0 1 0 112 0a56 56 0 1 0-112 0zm280 0a56 56 0 1 0 112 0a56 56 0 1 0-112 0z"/>`),
		g.Group(children),
	)
}