package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlusMinusBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m208.49 64.49l-144 144a12 12 0 0 1-17-17l144-144a12 12 0 0 1 17 17ZM60 112a12 12 0 0 0 24 0V84h28a12 12 0 0 0 0-24H84V32a12 12 0 0 0-24 0v28H32a12 12 0 0 0 0 24h28Zm164 60h-80a12 12 0 0 0 0 24h80a12 12 0 0 0 0-24Z"/>`),
		g.Group(children),
	)
}