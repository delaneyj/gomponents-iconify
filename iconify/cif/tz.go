package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<defs><path id="cifTz0" d="M.5.5h300v200H.5z"/></defs><defs><use href="#cifTz0" id="cifTz1"/></defs><g fill="none" fill-rule="evenodd"><path fill="#1EB53A" fill-rule="nonzero" d="M.5.5h300l-300 200"/><path fill="#00A3DD" fill-rule="nonzero" d="M.5 200.5h300V.5"/><mask id="cifTz2" fill="#fff"><use href="#cifTz0"/></mask><path stroke="#FCD116" stroke-width="70" d="m.5 200.5l300-200" mask="url(#cifTz2)"/><mask id="cifTz3" fill="#fff"><use href="#cifTz1"/></mask><path stroke="#000" stroke-width="50" d="m.5 200.5l300-200" mask="url(#cifTz3)"/></g>`),
		g.Group(children),
	)
}