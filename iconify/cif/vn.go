package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#DA251D" d="M.5.5h300v200H.5z"/><path fill="#FF0" d="m102.168 83.815l26.481 19.24l-10.115 31.131l26.482-19.239l26.481 19.239l-10.115-31.131l26.482-19.24h-32.841l-10.007-30.798l-10.007 30.798z"/></g>`),
		g.Group(children),
	)
}