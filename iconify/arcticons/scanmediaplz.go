package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scanmediaplz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.079 20.701a17.205 17.205 0 0 0-4.782-9.007a17.46 17.46 0 0 0-24.625 0M6.89 27.32a17.423 17.423 0 0 0 29.406 8.976"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10.121 10.149l3.458 3.455l4.621 4.617H5.5V5.5l4.621 4.649zm24.3 24.279l-4.653-4.649H42.5V42.5l-4.621-4.617l-3.458-3.455z"/>`),
		g.Group(children),
	)
}