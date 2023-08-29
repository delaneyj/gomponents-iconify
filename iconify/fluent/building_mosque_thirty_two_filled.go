package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingMosqueThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.25 6.25a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0ZM28 7.5A1.25 1.25 0 1 0 28 5a1.25 1.25 0 0 0 0 2.5ZM15.591 5.087a1 1 0 0 1 .818 0l.003.002a9.998 9.998 0 0 1 .348.164a24.551 24.551 0 0 1 3.576 2.133c.978.709 1.997 1.58 2.78 2.578c.778.994 1.384 2.194 1.384 3.536a8.47 8.47 0 0 1-.752 3.5H27v-7a1 1 0 1 1 2 0v17a2 2 0 0 1-2 2h-6a2 2 0 0 1-2-2v-2a3 3 0 1 0-6 0v2a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V10a1 1 0 0 1 2 0v7h3.252a8.47 8.47 0 0 1-.752-3.5c0-1.342.606-2.542 1.385-3.536c.782-.999 1.8-1.87 2.779-2.577a24.551 24.551 0 0 1 3.927-2.3Z"/>`),
		g.Group(children),
	)
}