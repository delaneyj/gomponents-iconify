package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProhibitThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 28a100 100 0 1 0 100 100A100.11 100.11 0 0 0 128 28Zm92 100a91.67 91.67 0 0 1-24.21 62.13L65.87 60.21A92 92 0 0 1 220 128Zm-184 0a91.67 91.67 0 0 1 24.21-62.13l129.92 129.92A92 92 0 0 1 36 128Z"/>`),
		g.Group(children),
	)
}