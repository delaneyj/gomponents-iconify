package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicRhythm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M5 42h5m-5-6h5m-5-6h5m-5-6h5m6 18h5m-5-6h5m-5-6h5m-5-6h5m-5-6h5m-5-6h5m-5-6h5m6 36h5m6 0h5m-16-6h5m6 0h5m-16-6h5m6 0h5m-5-6h5m-5-6h5"/>`),
		g.Group(children),
	)
}