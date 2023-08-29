package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackJo0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackJo3)"><use href="#flagpackJo0"/><path fill="#F7FCFF" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackJo1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g fill-rule="evenodd" clip-rule="evenodd" mask="url(#flagpackJo1)"><path fill="#272727" d="M0 0v8h32V0H0Z"/><path fill="#093" d="M0 16v8h32v-8H0Z"/></g><path fill="#E31D1C" fill-rule="evenodd" d="M0 0v24l20-12L0 0Z" clip-rule="evenodd"/><mask id="flagpackJo2" width="20" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24l20-12L0 0Z" clip-rule="evenodd"/></mask><g mask="url(#flagpackJo2)"><path fill="#F7FCFF" fill-rule="evenodd" d="M9.002 13.87L7.132 15l.426-2.204L6 11.146l2.11-.088L9.001 9l.892 2.058H12l-1.554 1.738l.468 2.204l-1.912-1.13Z" clip-rule="evenodd"/></g></g><defs><clipPath id="flagpackJo3"><use href="#flagpackJo0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}