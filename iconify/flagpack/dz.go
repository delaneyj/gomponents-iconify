package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackDz0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackDz1)"><use href="#flagpackDz0"/><path fill="#F7FCFF" fill-rule="evenodd" d="M16 0h16v24H16V0Z" clip-rule="evenodd"/><path fill="#36A400" fill-rule="evenodd" d="M0 0h16v24H0V0Z" clip-rule="evenodd"/><path fill="red" fill-rule="evenodd" d="M17.791 6.795c1.522 0 2.913.562 3.976 1.49a7.353 7.353 0 1 0 0 9.123a6.051 6.051 0 1 1-3.976-10.612Zm2.82 2.749l-1.703 1.929l-2.445-.674l1.356 2.12l-1.356 2.258l2.565-.924l1.402 2.189v-2.54l2.292-.982l-2.292-.821l.18-2.555Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackDz1"><use href="#flagpackDz0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}