package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpinnerGapBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M140 32v32a12 12 0 0 1-24 0V32a12 12 0 0 1 24 0Zm84 84h-32a12 12 0 0 0 0 24h32a12 12 0 0 0 0-24Zm-42.26 48.77a12 12 0 1 0-17 17l22.63 22.63a12 12 0 0 0 17-17ZM128 180a12 12 0 0 0-12 12v32a12 12 0 0 0 24 0v-32a12 12 0 0 0-12-12Zm-53.74-15.23L51.63 187.4a12 12 0 0 0 17 17l22.63-22.63a12 12 0 1 0-17-17ZM76 128a12 12 0 0 0-12-12H32a12 12 0 0 0 0 24h32a12 12 0 0 0 12-12Zm-7.4-76.37a12 12 0 1 0-17 17l22.66 22.6a12 12 0 0 0 17-17Z"/>`),
		g.Group(children),
	)
}