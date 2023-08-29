package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackGh0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackGh1)"><use href="#flagpackGh0"/><path fill="#5EAA22" fill-rule="evenodd" d="M0 16h32v8H0v-8Z" clip-rule="evenodd"/><path fill="#FECA00" fill-rule="evenodd" d="M0 8h32v8H0V8Z" clip-rule="evenodd"/><path fill="#E11C1B" fill-rule="evenodd" d="M0 0h32v8H0V0Z" clip-rule="evenodd"/><path fill="#1D1D1D" fill-rule="evenodd" d="m16.075 14.49l-3.485 2.418l1.114-4.14l-2.944-2.646l3.85-.143l1.465-4.095l1.466 4.095h3.827l-2.92 2.788l1.278 3.897l-3.65-2.174Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackGh1"><use href="#flagpackGh0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}