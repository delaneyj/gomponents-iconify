package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarRental(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M10.5 1a1 1 0 0 0-1 1H2v1l1 1l1-1l1 1l1-1l1 1h2.5a1 1 0 0 0 1 1h2a.5.5 0 0 0 .5-.5v-3a.5.5 0 0 0-.5-.5h-2Zm.5 1.5a.5.5 0 0 1 1 0v1a.5.5 0 0 1-1 0v-1ZM2.146 9.354A.5.5 0 0 0 2 9.707V13.5a.5.5 0 0 0 .5.5H4a.5.5 0 0 0 .5-.5V13h6v.5a.5.5 0 0 0 .5.5h1.5a.5.5 0 0 0 .5-.5V9.707a.5.5 0 0 0-.146-.353L12 8.5l-1.354-2.257a.5.5 0 0 0-.43-.243H4.784a.5.5 0 0 0-.429.243L3 8.5l-.854.854ZM11.134 9H3.866l1.2-2h4.868l1.2 2ZM5.5 10.828v.372a.3.3 0 0 1-.3.3H3.3a.3.3 0 0 1-.3-.3v-.834a.3.3 0 0 1 .359-.294l1.82.364a.4.4 0 0 1 .321.392Zm6.5-.34v.712a.3.3 0 0 1-.3.3H9.8a.3.3 0 0 1-.3-.3v-.454a.3.3 0 0 1 .241-.294l1.78-.356a.4.4 0 0 1 .479.392Z"/>`),
		g.Group(children),
	)
}