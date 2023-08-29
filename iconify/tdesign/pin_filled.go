package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18.076.981l4.949 4.95l-6.365 7.772l2.121 2.122l-5.305 5.305l-4.596-4.596l-6.718 6.718l-1.414-1.415l6.718-6.717l-4.597-4.596l5.306-5.306l2.121 2.122l7.78-6.36Z"/>`),
		g.Group(children),
	)
}