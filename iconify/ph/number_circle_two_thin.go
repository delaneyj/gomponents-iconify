package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleTwoThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 28a100 100 0 1 0 100 100A100.11 100.11 0 0 0 128 28Zm0 192a92 92 0 1 1 92-92a92.1 92.1 0 0 1-92 92Zm22.36-99.15L112 172h40a4 4 0 0 1 0 8h-48a4 4 0 0 1-3.2-6.4L144 116a20 20 0 0 0-4-28a20 20 0 0 0-28 4a20.08 20.08 0 0 0-2.89 5.37a4 4 0 0 1-7.55-2.66a28.19 28.19 0 0 1 4-7.52a28 28 0 1 1 44.72 33.7Z"/>`),
		g.Group(children),
	)
}