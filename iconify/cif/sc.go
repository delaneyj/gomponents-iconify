package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#D62828" d="M.5.5h300v150H.5z"/><path fill="#FCD856" d="M.5.5h200l-200 150z"/><path fill="#FFF" d="M300.5 50.5v100H.5z"/><path fill="#003F87" d="M.5.5h100l-100 150z"/><path fill="#007A3D" d="M.5 150.5h300v-50z"/></g>`),
		g.Group(children),
	)
}