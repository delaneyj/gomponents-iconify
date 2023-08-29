package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Poweroff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.074 2.328a.663.663 0 0 1-.143.927c-2.24 1.642-3.603 4.27-3.603 7.03c0 4.546 4.011 8.389 8.679 8.389c4.665 0 8.665-3.84 8.665-8.39c0-2.73-1.33-5.331-3.523-6.976a.663.663 0 1 1 .797-1.06C18.47 4.14 20 7.131 20 10.283C20 15.578 15.393 20 10.007 20C4.618 20 0 15.576 0 10.284c0-3.187 1.569-6.21 4.146-8.098a.664.664 0 0 1 .928.142ZM10 0a.7.7 0 0 1 .7.7v7.8a.7.7 0 0 1-1.4 0V.7A.7.7 0 0 1 10 0Z"/>`),
		g.Group(children),
	)
}