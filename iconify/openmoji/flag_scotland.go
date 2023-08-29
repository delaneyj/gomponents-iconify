package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagScotland(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#1e50a0" d="M5 17h62v38H5z"/><path fill="#fff" stroke="#fff" stroke-miterlimit="10" d="M10.6 54.72c-5.599.3-5.599.3-5.645-3.774l56.15-33.69c4.468.132 6.208-.809 5.843 3.674z"/><path fill="#fff" stroke="#fff" stroke-miterlimit="10" d="M10.3 17.03c-5.301-.033-5.301-.033-5.432 3.904l56.9 34.13c5.235-.065 5.235-.065 5.202-4.018z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}