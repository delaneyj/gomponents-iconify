package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnfoldLessDouble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8.85 23.95l-1.4-1.4l4.575-4.575L16.6 22.55l-1.4 1.425l-3.175-3.175l-3.175 3.15Zm0-5l-1.4-1.4l4.575-4.575L16.6 17.55l-1.4 1.425l-3.175-3.175l-3.175 3.15ZM12.025 11l-4.6-4.6L8.85 4.975l3.175 3.175l3.15-3.175L16.6 6.4L12.025 11Zm0-5l-4.6-4.6L8.85-.025l3.175 3.175l3.15-3.175L16.6 1.4L12.025 6Z"/>`),
		g.Group(children),
	)
}