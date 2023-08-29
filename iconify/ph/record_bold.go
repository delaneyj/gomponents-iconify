package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecordBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 20a108 108 0 1 0 108 108A108.12 108.12 0 0 0 128 20Zm0 192a84 84 0 1 1 84-84a84.09 84.09 0 0 1-84 84Zm0-152a68 68 0 1 0 68 68a68.07 68.07 0 0 0-68-68Zm0 112a44 44 0 1 1 44-44a44.05 44.05 0 0 1-44 44Z"/>`),
		g.Group(children),
	)
}