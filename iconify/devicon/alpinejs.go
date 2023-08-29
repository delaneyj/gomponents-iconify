package devicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alpinejs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#77c1d2" fill-rule="evenodd" d="M98.444 35.562L126 62.997L98.444 90.432L70.889 62.997z" clip-rule="evenodd"/><path fill="#2d3441" fill-rule="evenodd" d="m29.556 35.562l57.126 56.876H31.571L2 62.997z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}