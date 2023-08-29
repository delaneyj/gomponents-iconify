package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#6CF" d="M.5.5h300v150H.5z"/><path fill="#FFF" d="m100.5 137.5l50-7l50 7l-50-123.5z"/><path fill="#000" d="m107.25 137.5l43.25-7l43.25 7l-43.25-107z"/><path fill="#FCD116" d="M100.5 137.5h100l-50-62z"/></g>`),
		g.Group(children),
	)
}