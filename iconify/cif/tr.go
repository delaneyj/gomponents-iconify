package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#E30A17" d="M.5.5h300v200H.5z"/><circle cx="106.75" cy="100.5" r="50" fill="#FFF"/><circle cx="119.25" cy="100.5" r="40" fill="#E30A17"/><path fill="#FFF" d="m146.334 100.5l45.225 14.695l-27.951-38.472v47.553l27.951-38.471z"/></g>`),
		g.Group(children),
	)
}