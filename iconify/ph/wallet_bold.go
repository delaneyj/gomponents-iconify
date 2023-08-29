package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WalletBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M196 140a16 16 0 1 1-16-16a16 16 0 0 1 16 16Zm40-32v80a32 32 0 0 1-32 32H60a32 32 0 0 1-32-32V68.92A32 32 0 0 1 60 36h132a12 12 0 0 1 0 24H60a8 8 0 0 0-8 8.26v.08A8.32 8.32 0 0 0 60.48 76H204a32 32 0 0 1 32 32Zm-24 0a8 8 0 0 0-8-8H60.48A33.72 33.72 0 0 1 52 98.92V188a8 8 0 0 0 8 8h144a8 8 0 0 0 8-8Z"/>`),
		g.Group(children),
	)
}