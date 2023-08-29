package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackCw0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackCw2)"><use href="#flagpackCw0"/><path fill="#2E42A5" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackCw1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g fill-rule="evenodd" clip-rule="evenodd" mask="url(#flagpackCw1)"><path fill="#F7FCFF" d="M4.254 6.15L2.266 7.198l.96-1.868L2 3.963l1.523-.056L4.254 2l.558 1.907l1.783.056l-1.284 1.368l.815 1.868l-1.872-1.05Zm7.098 4.926l-2.454 1.029l.963-2.577L7.726 7.8h2.558l1.068-2.787l.816 2.787h2.56l-1.82 1.73l.905 2.576l-2.46-1.029Z"/><path fill="#F9E813" d="M0 14v4h32v-4H0Z"/></g></g><defs><clipPath id="flagpackCw2"><use href="#flagpackCw0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}