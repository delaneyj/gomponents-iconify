package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 28a100 100 0 1 0 100 100A100.11 100.11 0 0 0 128 28ZM68.87 198.42a68 68 0 0 1 118.26 0a91.8 91.8 0 0 1-118.26 0Zm124.3-5.55a75.61 75.61 0 0 0-44.51-34a44 44 0 1 0-41.32 0a75.61 75.61 0 0 0-44.51 34a92 92 0 1 1 130.34 0ZM128 156a36 36 0 1 1 36-36a36 36 0 0 1-36 36Z"/>`),
		g.Group(children),
	)
}