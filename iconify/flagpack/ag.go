package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackAg0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackAg1)"><use href="#flagpackAg0"/><path fill="#1B1B1B" fill-rule="evenodd" d="M0 0h32v24H0V0Z" clip-rule="evenodd"/><path fill="#F9D313" fill-rule="evenodd" d="m16 14l-2.003 3.8l.05-4.196l-3.658 2.65l2.093-3.76l-4.59.977l3.72-2.58L7 10l4.613-.89L7.89 6.529l4.59.977l-2.092-3.76l3.659 2.65l-.05-4.195L16 6l2.003-3.8l-.05 4.196l3.658-2.65l-2.093 3.76l4.59-.977l-3.72 2.58L25 10l-4.613.89l3.722 2.581l-4.59-.977l2.092 3.76l-3.659-2.65l.05 4.195L16 14Z" clip-rule="evenodd"/><path fill="#F1F9FF" d="M6 14h20v10H6z"/><path fill="#4A80E8" d="M2 10h28v4H2z"/><path fill="#E31D1C" fill-rule="evenodd" d="m0 6l16 18L32 6v18H0V6Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackAg1"><use href="#flagpackAg0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}