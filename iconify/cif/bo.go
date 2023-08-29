package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#007934" d="M.5.227h300v204.545H.5z"/><path fill="#F9E300" d="M.5.227h300v136.364H.5z"/><path fill="#D52B1E" d="M.5.227h300v68.182H.5z"/></g>`),
		g.Group(children),
	)
}