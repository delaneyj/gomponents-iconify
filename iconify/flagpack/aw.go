package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackAw0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackAw2)"><use href="#flagpackAw0"/><path fill="#5BA3DA" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackAw1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g fill-rule="evenodd" clip-rule="evenodd" mask="url(#flagpackAw1)"><path fill="red" stroke="#F7FCFF" stroke-width=".35" d="m5.773 7.96l-4.38-.925l4.434-.878l1.283-4.59L8.122 6.12l3.96.92l-3.914.92l-1.129 3.743L5.773 7.96Z"/><path fill="#FAD615" d="M32 14H0v2h32v-2Zm0 4H0v2h32v-2Z"/></g></g><defs><clipPath id="flagpackAw2"><use href="#flagpackAw0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}