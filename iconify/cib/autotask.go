package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Autotask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m23.907 15.5l3.708 9.317h-5.729l-2.432-6.312l-13.183 8.891l-4.937-10.593h1.088l5.849 4.557l10.151-5.532l-1.448-3.76l-1.557 3.307l-6.583 3.62l4.735-12.453h6.771l2.697 6.771l8.964-4.885V.001h-32v32h32V10.042z"/>`),
		g.Group(children),
	)
}