package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reason(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M0 0v32h32V0zm15.359 29.188l-1.771-3.385h-2.391v3.385H7.999V16.922h5.505c3.255 0 5.109 1.583 5.109 4.318c0 1.854-.776 3.198-2.188 3.948l2.557 4zm15.146 0h-9.839V16.922h9.651v2.557h-6.453v2.292h5.844v2.536l-5.844.01v2.307h6.656v2.563zm-14.932-7.855c0 1.24-.745 1.906-2.042 1.906h-2.333v-3.76h2.333c1.297 0 2.042.656 2.042 1.87z"/>`),
		g.Group(children),
	)
}