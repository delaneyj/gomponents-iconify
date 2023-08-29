package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#CE1126" d="M.5.5h100v150H.5z"/><path fill="#FCD116" d="M100.5.5h200v75h-200z"/><path fill="#009E49" d="M100.5 75.5h200v75h-200z"/><path fill="#000" d="m26.724 67.774l14.694 10.677l-5.612 17.274L50.5 85.049l14.695 10.676l-5.613-17.274l14.695-10.677H56.113L50.5 50.5l-5.613 17.274z"/></g>`),
		g.Group(children),
	)
}