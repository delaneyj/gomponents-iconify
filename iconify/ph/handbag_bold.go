package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandbagBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m243.86 197.65l-14.25-120A20.06 20.06 0 0 0 209.67 60h-29.84a52 52 0 0 0-103.66 0H46.33a20.06 20.06 0 0 0-19.94 17.65l-14.25 120A20 20 0 0 0 32.08 220h191.84a20 20 0 0 0 19.94-22.35ZM128 36a28 28 0 0 1 27.71 24h-55.42A28 28 0 0 1 128 36ZM36.5 196L49.81 84H76v20a12 12 0 0 0 24 0V84h56v20a12 12 0 0 0 24 0V84h26.19l13.31 112Z"/>`),
		g.Group(children),
	)
}