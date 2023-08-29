package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackLr0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackLr1)"><use href="#flagpackLr0"/><path fill="#F7FCFF" fill-rule="evenodd" d="M0 0h32v24H0V0Z" clip-rule="evenodd"/><path fill="#E31D1C" d="M.027 5.5h32v3h-32z"/><path fill="#E31D1C" fill-rule="evenodd" d="M0 0h32v3H0V0Z" clip-rule="evenodd"/><path fill="#E31D1C" d="M-.059 11h32v3h-32zm.171 5.4h32v3h-32zm-.01 5.1h32v3h-32z"/><path fill="#3D58DB" fill-rule="evenodd" d="M0 0h16v14H0V0Z" clip-rule="evenodd"/><path fill="#F7FCFF" fill-rule="evenodd" d="m8.132 9.213l-2.92 2.026l.933-3.47L4 5.552l2.904-.12L8.132 2l1.23 3.432h2.898l-2.14 2.337l1.072 3.266l-3.06-1.822Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackLr1"><use href="#flagpackLr0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}