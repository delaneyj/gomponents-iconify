package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackSc0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackSc2)"><use href="#flagpackSc0"/><path fill="#2E42A5" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackSc1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g mask="url(#flagpackSc1)"><path fill="#FFD018" fill-rule="evenodd" d="M0 23.997L16.151-2h16.151L0 23.997Z" clip-rule="evenodd"/><path fill="#E31D1C" d="m0 23.997l34.463-12.999V-5.103L0 23.997Z"/><path fill="#F7FCFF" d="m0 23.997l34.463-4.999v-8.101L0 23.997Z"/><path fill="#5EAA22" fill-rule="evenodd" d="M0 23.997h34.463v-7.1L0 23.997Z" clip-rule="evenodd"/></g></g><defs><clipPath id="flagpackSc2"><use href="#flagpackSc0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}