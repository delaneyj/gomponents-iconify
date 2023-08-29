package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#024FA2" d="M.5.5h300v150H.5z"/><path fill="#FFF" d="M.5 25.5h300v100H.5z"/><path fill="#ED1C27" d="M.5 29.666h300v91.667H.5z"/><circle cx="100.5" cy="75.5" r="33.333" fill="#FFF"/><path fill="#ED1C27" d="m69.789 65.521l18.98 13.79l-7.249 22.313l18.98-13.79l18.981 13.79l-7.25-22.313l18.98-13.79H107.75l-7.25-22.313l-7.25 22.313z"/></g>`),
		g.Group(children),
	)
}