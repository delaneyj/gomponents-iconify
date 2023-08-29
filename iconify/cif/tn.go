package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#E70013" d="M.5.5h300v200H.5z"/><circle cx="150.5" cy="100.5" r="50" fill="#FFF"/><circle cx="150.5" cy="100.5" r="37.5" fill="#E70013"/><circle cx="160.5" cy="100.5" r="30" fill="#FFF"/><path fill="#E70013" d="m138 100.5l40.703-13.225l-25.156 34.624V79.102l25.156 34.624z"/></g>`),
		g.Group(children),
	)
}