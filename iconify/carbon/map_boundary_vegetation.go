package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapBoundaryVegetation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 16H18v-2h5v-4a2 2 0 0 1 2-2h1.5a1.5 1.5 0 0 0 .284-2.973l-.711-.136l-.075-.571a2.632 2.632 0 0 0-1.68-2.182a2.505 2.505 0 0 0-3.083 1.304l-.304.647l-.71-.077a2.018 2.018 0 0 0-2.17 2.434A2.098 2.098 0 0 0 20.126 8H21v2h-.782a4.15 4.15 0 0 1-4.206-3.686a4.005 4.005 0 0 1 3.741-4.306a4.5 4.5 0 0 1 8.074 1.252a3.505 3.505 0 0 1 2.142 3.71A3.628 3.628 0 0 1 26.317 10H25v4h3zm-10 2h4v2h-4zm-4 0v-2h-2v4h4v-2h-2zm-2-8h2v4h-2z"/><path fill="currentColor" d="m17.885 30l-6.066-3.858L4 27.989V4h10v4h-2V6H6v19.461l6.181-1.46l5.934 3.775L22.818 26H26v-6h-2v-2h4v10h-4.818l-5.297 2z"/>`),
		g.Group(children),
	)
}