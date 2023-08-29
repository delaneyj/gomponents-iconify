package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoadSpikes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M64 116.8c0-15.8 20.5-22 29.3-8.9L192 256V116.8c0-15.8 20.5-22 29.3-8.9L320 256V116.8c0-15.8 20.5-22 29.3-8.9L448 256V116.8c0-15.8 20.5-22 29.3-8.9l129.5 194.3c14.2 21.3-1.1 49.7-26.6 49.7H64V116.8zM32 384h576c17.7 0 32 14.3 32 32s-14.3 32-32 32H32c-17.7 0-32-14.3-32-32s14.3-32 32-32z"/>`),
		g.Group(children),
	)
}