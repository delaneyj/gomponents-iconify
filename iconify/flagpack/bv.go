package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackBv0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackBv2)"><use href="#flagpackBv0"/><path fill="#E31D1C" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackBv1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g mask="url(#flagpackBv1)"><path fill="#2E42A5" stroke="#F7FCFF" stroke-width="2" d="M10-1H9V9H-1v6H9v10h6V15h18V9H15V-1h-5Z"/></g></g><defs><clipPath id="flagpackBv2"><use href="#flagpackBv0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}