package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#009543" d="M.5.5h300v200H.5z"/><path fill="#FBDE4A" d="m.5 200.5l200-200h100v200z"/><path fill="#DC241F" d="M300.5.5v200h-200z"/></g>`),
		g.Group(children),
	)
}