package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sharpness(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.5.586V19.5H1.587L20.501.586ZM6.416 17.5H18.5V5.414L6.415 17.5ZM20.5 21v2h-19v-2h19Z"/>`),
		g.Group(children),
	)
}