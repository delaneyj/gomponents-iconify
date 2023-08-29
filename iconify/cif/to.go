package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func To(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#C10000" d="M.5.5h300v150H.5z"/><path fill="#FFF" d="M.5.5h125v75H.5z"/><path fill="#C10000" d="M53.625 9.875h18.75v56.25h-18.75z"/><path fill="#C10000" d="M34.875 28.625h56.25v18.75h-56.25z"/></g>`),
		g.Group(children),
	)
}