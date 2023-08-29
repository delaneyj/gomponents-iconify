package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TennisBallThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M198.74 57.29A99.39 99.39 0 0 0 128 28h-.19a100 100 0 1 0 70.93 29.29ZM62.92 63a91.38 91.38 0 0 1 60.8-26.85a91.31 91.31 0 0 1-26.85 60.74a91.44 91.44 0 0 1-60.77 26.84A91.61 91.61 0 0 1 62.92 63Zm-26.83 68.74a99.36 99.36 0 0 0 66.44-29.2a99.25 99.25 0 0 0 29.21-66.45a91.9 91.9 0 0 1 88.17 88.17a100 100 0 0 0-95.65 95.65a91.9 91.9 0 0 1-88.17-88.17Zm157 61.31a91.38 91.38 0 0 1-60.8 26.85a92 92 0 0 1 87.62-87.63a91.61 91.61 0 0 1-26.83 60.78Z"/>`),
		g.Group(children),
	)
}