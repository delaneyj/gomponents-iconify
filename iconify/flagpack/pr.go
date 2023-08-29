package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackPr0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackPr3)"><use href="#flagpackPr0"/><path fill="#EF0000" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackPr1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g mask="url(#flagpackPr1)"><path fill="#EF0000" stroke="#F7FCFF" stroke-width="4" d="M0 8h-2v8h36V8H0Z"/></g><path fill="#3D58DB" fill-rule="evenodd" d="M0 0v24l18-12L0 0Z" clip-rule="evenodd"/><mask id="flagpackPr2" width="18" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24l18-12L0 0Z" clip-rule="evenodd"/></mask><g mask="url(#flagpackPr2)"><path fill="#F7FCFF" fill-rule="evenodd" d="m6.688 14.217l-3.672 1.938l1.787-3.894l-2.277-2.08l2.812-.104l1.35-3.52l1.03 3.52h2.807l-1.87 2.184l1.488 3.894l-3.455-1.938Z" clip-rule="evenodd"/></g></g><defs><clipPath id="flagpackPr3"><use href="#flagpackPr0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}