package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightningLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M213.84 118.63a6 6 0 0 0-3.73-4.25l-59.23-22.21l15-75a6 6 0 0 0-10.27-5.27l-112 120a6 6 0 0 0 2.28 9.71l59.23 22.21l-15 75a6 6 0 0 0 3.14 6.52A6.07 6.07 0 0 0 96 246a6 6 0 0 0 4.39-1.91l112-120a6 6 0 0 0 1.45-5.46ZM106 220.46l11.85-59.28a6 6 0 0 0-3.77-6.8l-55.6-20.85l91.46-98l-11.82 59.29a6 6 0 0 0 3.77 6.8l55.6 20.85Z"/>`),
		g.Group(children),
	)
}