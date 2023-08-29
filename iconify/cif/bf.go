package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#009E49" d="M.5.5h300v200H.5z"/><path fill="#EF2B2D" d="M.5.5h300v100H.5z"/><path fill="#FCD116" d="m118.798 90.199l19.593 14.235l-7.484 23.033l19.593-14.235l19.593 14.235l-7.484-23.033l19.593-14.235h-24.218L150.5 67.166l-7.374 23.033z"/></g>`),
		g.Group(children),
	)
}