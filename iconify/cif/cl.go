package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#D52B1E" d="M300.5 200.5H.5V.5h300z"/><path fill="#FFF" d="M100.5 100.5h200V.5H.5z"/><path fill="#0039A6" d="M100.5 100.5H.5V.5h100z"/><path fill="#FFF" d="m26.724 42.774l14.694 10.677l-5.612 17.274L50.5 60.049l14.695 10.676l-5.613-17.274l14.695-10.677H56.113L50.5 25.5l-5.613 17.274z"/></g>`),
		g.Group(children),
	)
}