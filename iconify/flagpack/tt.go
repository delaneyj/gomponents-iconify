package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackTt0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackTt2)"><use href="#flagpackTt0"/><path fill="#E31D1C" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackTt1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g mask="url(#flagpackTt1)"><path fill="#272727" stroke="#F7FCFF" stroke-width="1.5" d="m29.56 29.496l-.543.444l-.48-.512L-1.807-2.971l-.548-.585l.62-.508l3.097-2.532l.543-.444l.48.512L32.727 25.87l.549.586l-.621.508l-3.096 2.532Z"/></g></g><defs><clipPath id="flagpackTt2"><use href="#flagpackTt0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}