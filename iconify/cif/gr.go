package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><path fill="#0D5EAF" fill-rule="nonzero" d="M.5.5h300v200H.5z"/><path stroke="#FFF" stroke-width="20" d="M56.056.5v122.222M.5 56.055h111.111m0-22.222H300.5M111.611 78.278H300.5M.5 122.722h300M.5 167.167h300"/></g>`),
		g.Group(children),
	)
}