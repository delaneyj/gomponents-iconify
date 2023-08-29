package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackMv0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackMv2)"><use href="#flagpackMv0"/><path fill="#C51918" fill-rule="evenodd" d="M0 0h32v22a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V0Z" clip-rule="evenodd"/><path fill="#C51918" fill-rule="evenodd" d="M0 0h32v24H0V0Z" clip-rule="evenodd"/><path fill="#579D20" fill-rule="evenodd" d="M6 6h20v12H6V6Z" clip-rule="evenodd"/><path stroke="#B6EB9A" stroke-opacity=".238" stroke-width="2" d="M7 7h18v10H7V7Z"/><mask id="flagpackMv1" width="20" height="12" x="6" y="6" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M6 6h20v12H6V6Z" clip-rule="evenodd"/><path stroke="#fff" stroke-width="2" d="M7 7h18v10H7V7Z"/></mask><g mask="url(#flagpackMv1)"><path fill="#F9FAFA" fill-rule="evenodd" d="M16.033 12.463c-.017 3.065 2.396 4.7 2.396 4.7c-2.753.323-4.586-2.174-4.586-4.67c0-2.498 2.498-4.56 4.586-5.492c0 0-2.379 2.396-2.396 5.462Z" clip-rule="evenodd"/></g></g><defs><clipPath id="flagpackMv2"><use href="#flagpackMv0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}