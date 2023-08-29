package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackSn0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackSn1)"><use href="#flagpackSn0"/><path fill="#FBCD17" fill-rule="evenodd" d="M10 0h12v24H10V0Z" clip-rule="evenodd"/><path fill="#006923" fill-rule="evenodd" d="m16.075 14.49l-3.485 2.418l1.114-4.14l-2.56-2.79h3.466l1.465-4.094l1.466 4.095h3.46l-2.554 2.788l1.279 3.897l-3.65-2.174Z" clip-rule="evenodd"/><path fill="#E11C1B" fill-rule="evenodd" d="M22 0h10v24H22V0Z" clip-rule="evenodd"/><path fill="#006923" fill-rule="evenodd" d="M0 0h10v24H0V0Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackSn1"><use href="#flagpackSn0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}