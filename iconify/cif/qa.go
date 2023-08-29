package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Qa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#8D1B3D" d="M.5.445h300v118.11H.5z"/><path fill="#FFF" d="M.5.445v118.11h94.488l24.003-6.558l-24.003-6.57l24.003-6.558l-24.003-6.558l24.003-6.57l-24.003-6.558l24.003-6.558l-24.003-6.57l24.003-6.558l-24.003-6.558l24.003-6.57l-24.003-6.558l24.003-6.558l-24.003-6.57l24.003-6.558l-24.003-6.558l24.003-6.57L94.988.439H.5z"/></g>`),
		g.Group(children),
	)
}