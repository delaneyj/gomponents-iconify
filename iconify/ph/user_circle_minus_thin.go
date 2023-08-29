package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserCircleMinusThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M172 56a4 4 0 0 1 4-4h48a4 4 0 0 1 0 8h-48a4 4 0 0 1-4-4Zm54.62 55.34a99.89 99.89 0 1 1-82-82a4 4 0 0 1-1.32 7.89A93.4 93.4 0 0 0 128 36a92 92 0 0 0-65.17 156.87a75.61 75.61 0 0 1 44.51-34a44 44 0 1 1 41.32 0a75.61 75.61 0 0 1 44.51 34A91.69 91.69 0 0 0 220 128a93.58 93.58 0 0 0-1.27-15.34a4 4 0 0 1 7.89-1.32ZM128 156a36 36 0 1 0-36-36a36 36 0 0 0 36 36Zm0 64a91.61 91.61 0 0 0 59.14-21.58a68 68 0 0 0-118.27 0A91.56 91.56 0 0 0 128 220Z"/>`),
		g.Group(children),
	)
}