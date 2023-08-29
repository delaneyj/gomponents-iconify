package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackGy0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackGy1)"><use href="#flagpackGy0"/><path fill="#5EAA22" fill-rule="evenodd" d="M0 0h32v24H0V0Z" clip-rule="evenodd"/><path fill="#FECA00" stroke="#F7FCFF" stroke-width="2" d="M1 22.587V1.413L30.995 12L1 22.587Z"/><path fill="#E11C1B" stroke="#272727" stroke-width="2" d="M-1 23.955V.045L14.371 12L-1 23.955Z"/></g><defs><clipPath id="flagpackGy1"><use href="#flagpackGy0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}