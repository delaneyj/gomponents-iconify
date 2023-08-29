package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackLv0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackLv1)"><use href="#flagpackLv0"/><path fill="#C51918" fill-rule="evenodd" d="M0 14h32v10H0V14Z" clip-rule="evenodd"/><path fill="#F7FCFF" fill-rule="evenodd" d="M0 8h32v6H0V8Z" clip-rule="evenodd"/><path fill="#C51918" fill-rule="evenodd" d="M0 0h32v10H0V0Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackLv1"><use href="#flagpackLv0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}