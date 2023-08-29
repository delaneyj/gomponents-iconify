package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PointOut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPointOut0"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M37 44H17.476a.257.257 0 0 1-.218-.121L7.86 28.727a4.095 4.095 0 1 1 7.011-4.23l2.462 4.194V7.942a3.942 3.942 0 0 1 7.884 0v9.329c0 .585.465 1.066 1.05 1.085l11.621.388A2.185 2.185 0 0 1 40 20.928V41a3 3 0 0 1-3 3Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPointOut0)"/>`),
		g.Group(children),
	)
}