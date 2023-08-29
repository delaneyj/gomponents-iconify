package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#00853F" d="M.5.5h100v200H.5z"/><path fill="#FDEF42" d="M100.5.5h100v200h-100z"/><path fill="#E31B23" d="M200.5.5h100v200h-100z"/><path fill="#00853F" d="m118.798 90.199l19.593 14.235l-7.484 23.033l19.593-14.235l19.593 14.235l-7.484-23.033l19.593-14.235h-24.218L150.5 67.166l-7.484 23.033z"/></g>`),
		g.Group(children),
	)
}