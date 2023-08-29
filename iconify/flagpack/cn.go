package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackCn0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackCn1)"><use href="#flagpackCn0"/><path fill="#E31D1C" fill-rule="evenodd" d="M0 0h32v24H0V0Z" clip-rule="evenodd"/><path fill="#FECA00" fill-rule="evenodd" d="m15.016 4.548l-1.01.61l.23-1.19l-.841-.89l1.139-.049l.482-1.11l.482 1.11h1.137l-.84.94l.253 1.19l-1.032-.61ZM7.018 9.607l-2.881 1.551l.657-3.026l-2.4-2.265l3.25-.123l1.374-2.826l1.374 2.826h3.243L9.24 8.132l.72 3.026l-2.943-1.55Zm9.998-1.059l-1.01.61l.23-1.19l-.841-.89l1.139-.049l.482-1.11l.482 1.11h1.137l-.84.94l.253 1.19l-1.032-.61Zm-1 4l-1.01.61l.23-1.19l-.841-.89l1.139-.049l.482-1.11l.482 1.11h1.137l-.84.94l.253 1.19l-1.032-.61Zm-3 3l-1.01.61l.23-1.19l-.841-.89l1.139-.049l.482-1.11l.482 1.11h1.137l-.84.94l.253 1.19l-1.032-.61Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackCn1"><use href="#flagpackCn0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}