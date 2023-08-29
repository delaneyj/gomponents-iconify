package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#FFF" d="M.5.833h300v183.333H.5z"/><path fill="#003580" d="M.5 67.5h300v50H.5z"/><path fill="#003580" d="M83.834.833h50v183.333h-50z"/></g>`),
		g.Group(children),
	)
}