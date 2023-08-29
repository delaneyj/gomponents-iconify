package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecordThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 28a100 100 0 1 0 100 100A100.11 100.11 0 0 0 128 28Zm0 192a92 92 0 1 1 92-92a92.1 92.1 0 0 1-92 92Zm0-160a68 68 0 1 0 68 68a68.07 68.07 0 0 0-68-68Zm0 128a60 60 0 1 1 60-60a60.07 60.07 0 0 1-60 60Z"/>`),
		g.Group(children),
	)
}