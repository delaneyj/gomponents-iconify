package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ua(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackUa0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackUa2)"><use href="#flagpackUa0"/><path fill="#3195F9" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackUa1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g mask="url(#flagpackUa1)"><path fill="#FECA00" fill-rule="evenodd" d="M0 12v12h32V12H0Z" clip-rule="evenodd"/></g></g><defs><clipPath id="flagpackUa2"><use href="#flagpackUa0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}