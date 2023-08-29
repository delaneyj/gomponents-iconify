package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackSr0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackSr1)"><use href="#flagpackSr0"/><path fill="#4E8B1D" fill-rule="evenodd" d="M0 16h32v8H0v-8ZM0 0h32v6H0V0Z" clip-rule="evenodd"/><path fill="#AF0100" stroke="#fff" stroke-width="3" d="M0 6.5h-1.5v11h35v-11H0Z"/><path fill="#FECA00" fill-rule="evenodd" d="M16.002 14.494L13.508 16l.57-2.938l-2.078-2.2l2.813-.118L16.003 8l1.19 2.744H20l-2.072 2.318l.624 2.938l-2.55-1.506Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackSr1"><use href="#flagpackSr0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}