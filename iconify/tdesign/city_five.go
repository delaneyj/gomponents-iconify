package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CityFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.5 1.72L15 4.52v5.108a4.124 4.124 0 0 1-6 0V4.519l-3.5-2.8L2 4.52V22h20V4.52l-3.5-2.8ZM9 12.14a6.151 6.151 0 0 0 6 0V20h-2v-5h-2v5H9v-7.86ZM7 20H4V5.48l1.5-1.2L7 5.48V20Zm10 0V5.48l1.5-1.2l1.5 1.2V20h-3Z"/>`),
		g.Group(children),
	)
}