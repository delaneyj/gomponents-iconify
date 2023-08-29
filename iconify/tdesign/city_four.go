package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CityFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.5 1.72L15 4.52v3.862l-3-1.5l-3 1.5V4.519l-3.5-2.8L2 4.52V22h20V4.52l-3.5-2.8ZM9 10.617l3-1.5l3 1.5V20h-2v-4h-2v4H9v-9.382ZM7 20H4V5.48l1.5-1.2L7 5.48V20Zm10 0V5.48l1.5-1.2l1.5 1.2V20h-3Zm-3-8h-4v2h4v-2Z"/>`),
		g.Group(children),
	)
}