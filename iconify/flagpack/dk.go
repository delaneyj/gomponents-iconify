package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackDk0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackDk2)"><use href="#flagpackDk0"/><path fill="#C51918" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackDk1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g mask="url(#flagpackDk1)"><path fill="#F7FCFF" fill-rule="evenodd" d="M10 0h4v10h18v4H14v10h-4V14H0v-4h10V0Z" clip-rule="evenodd"/></g></g><defs><clipPath id="flagpackDk2"><use href="#flagpackDk0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}