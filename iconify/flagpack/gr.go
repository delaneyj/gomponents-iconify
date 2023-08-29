package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackGr0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackGr1)"><use href="#flagpackGr0"/><path fill="#F7FCFF" fill-rule="evenodd" d="M0 0h32v24H0V0Z" clip-rule="evenodd"/><path fill="#4564F9" fill-rule="evenodd" d="M0 5.5h32v2.957H0V5.5Zm0 5.315h32v2.957H0v-2.957Zm32 5.128H0V18.9h32v-2.957ZM0 0h32v3H0V0Z" clip-rule="evenodd"/><path fill="#4564F9" d="M0 21h32v3H0z"/><path fill="#4564F9" fill-rule="evenodd" d="M0 0h15v13.8H0V0Z" clip-rule="evenodd"/><path fill="#F7FCFF" fill-rule="evenodd" d="M6 0h3v5.504h6v2.935H9V14.5H6V8.439H0V5.504h6V0Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackGr1"><use href="#flagpackGr0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}