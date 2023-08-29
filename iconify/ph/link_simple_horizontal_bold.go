package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkSimpleHorizontalBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M80 116h96a12 12 0 0 1 0 24H80a12 12 0 0 1 0-24Zm24 48H64a36 36 0 0 1 0-72h40a12 12 0 0 0 0-24H64a60 60 0 0 0 0 120h40a12 12 0 0 0 0-24Zm88-96h-40a12 12 0 0 0 0 24h40a36 36 0 0 1 0 72h-40a12 12 0 0 0 0 24h40a60 60 0 0 0 0-120Z"/>`),
		g.Group(children),
	)
}