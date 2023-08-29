package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func St(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#12AD2B" d="M.5.5h300v150H.5z"/><path fill="#FFCE00" d="M.5 43.357h300v64.286H.5z"/><path fill="#D21034" d="M.5.5v150l75-75"/><path fill="#000" d="m130.12 68.878l12.596 9.151l-4.811 14.807l12.595-9.151l12.596 9.151l-4.811-14.807l12.595-9.151h-15.569L150.5 54.071l-4.811 14.807zm75 0l12.596 9.151l-4.811 14.807l12.595-9.151l12.596 9.151l-4.811-14.807l12.595-9.151h-15.569L225.5 54.071l-4.811 14.807z"/></g>`),
		g.Group(children),
	)
}