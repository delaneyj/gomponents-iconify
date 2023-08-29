package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zoysii(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 4l18.67 14.81L24 33.62L5.33 18.81Zm12.13 20l6.54 5.19L24 44L5.33 29.19L11.87 24"/>`),
		g.Group(children),
	)
}