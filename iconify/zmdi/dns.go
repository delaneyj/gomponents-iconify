package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dns(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M363 213q8 0 14.5 6.5T384 235v128q0 8-6.5 14.5T363 384H21q-8 0-14.5-6.5T0 363V235q0-9 6.5-15.5T21 213h342zM85.5 341q17.5 0 30-12.5t12.5-30t-12.5-30t-30-12.5t-30 12.5t-12.5 30t12.5 30t30 12.5zM363 0q8 0 14.5 6.5T384 21v128q0 9-6.5 15.5T363 171H21q-8 0-14.5-6.5T0 149V21q0-8 6.5-14.5T21 0h342zM85.5 128q17.5 0 30-12.5t12.5-30t-12.5-30t-30-12.5t-30 12.5t-12.5 30t12.5 30t30 12.5z"/>`),
		g.Group(children),
	)
}