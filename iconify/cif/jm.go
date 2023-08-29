package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<defs><path id="cifJm0" d="M.5.5h300v150H.5z"/></defs><g fill="none" fill-rule="evenodd"><path fill="#009B3A" fill-rule="nonzero" d="M.5.5h300v150H.5z"/><mask id="cifJm1" fill="#fff"><use href="#cifJm0"/></mask><path fill="#000" fill-rule="nonzero" stroke="#FED100" stroke-width="20" d="m300.5.5l-150 75l150 75M.5.5l150 75l-150 75" mask="url(#cifJm1)"/></g>`),
		g.Group(children),
	)
}