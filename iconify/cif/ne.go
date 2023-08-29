package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#0DB02B" d="M.5.928h300v257.143H.5z"/><path fill="#FFF" d="M.5.928h300v171.429H.5z"/><path fill="#E05206" d="M.5.928h300v85.714H.5z"/><circle cx="150.5" cy="129.5" r="36.429" fill="#E05206"/></g>`),
		g.Group(children),
	)
}