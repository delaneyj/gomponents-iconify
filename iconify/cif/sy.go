package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#000" d="M.5.5h300v200H.5z"/><path fill="#FFF" d="M.5.5h300v133.333H.5z"/><path fill="#CE1126" d="M.5.5h300v66.667H.5z"/><path fill="#007A3D" d="m189.894 93.804l12.683 9.215l-4.812 15.009l12.735-9.252l12.735 9.252l-4.864-14.971l12.735-9.253h-15.741L210.5 78.833l-4.864 14.971zm-120 0l12.683 9.215l-4.812 15.009l12.735-9.252l12.735 9.252l-4.864-14.971l12.735-9.253H95.365L90.5 78.833l-4.864 14.971z"/></g>`),
		g.Group(children),
	)
}