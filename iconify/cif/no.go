package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func No(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#EF2B2D" d="M.5.409h300v218.182H.5z"/><path fill="#FFF" d="M82.318.409h54.545v218.182H82.318z"/><path fill="#FFF" d="M.5 82.227h300v54.545H.5z"/><path fill="#002868" d="M95.955.409h27.273v218.182H95.955z"/><path fill="#002868" d="M.5 95.863h300v27.273H.5z"/></g>`),
		g.Group(children),
	)
}