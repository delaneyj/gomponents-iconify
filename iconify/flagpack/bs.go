package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackBs0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackBs1)"><use href="#flagpackBs0"/><path fill="#FECA00" fill-rule="evenodd" d="M0 0h32v24H0V0Z" clip-rule="evenodd"/><path fill="#3CB1CF" fill-rule="evenodd" d="M0 0v8h32V0H0Zm0 16v8h32v-8H0Z" clip-rule="evenodd"/><path fill="#272727" fill-rule="evenodd" d="m0 0l16 12L0 24V0Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackBs1"><use href="#flagpackBs0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}