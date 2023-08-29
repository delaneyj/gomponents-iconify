package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CopyrightThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 28a100 100 0 1 0 100 100A100.11 100.11 0 0 0 128 28Zm0 192a92 92 0 1 1 92-92a92.1 92.1 0 0 1-92 92Zm35.2-65.59a44 44 0 1 1 0-52.82a4 4 0 0 1-6.4 4.81a36 36 0 1 0 0 43.2a4 4 0 0 1 6.4 4.81Z"/>`),
		g.Group(children),
	)
}