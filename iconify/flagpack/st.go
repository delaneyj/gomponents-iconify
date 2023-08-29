package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func St(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackSt0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackSt1)"><use href="#flagpackSt0"/><path fill="#FBCD17" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><path fill="#73BE4A" fill-rule="evenodd" d="M0 0v8h32V0H0Zm0 16v8h32v-8H0Z" clip-rule="evenodd"/><path fill="#C51918" fill-rule="evenodd" d="M0 0v24l10-12L0 0Z" clip-rule="evenodd"/><path fill="#272727" fill-rule="evenodd" d="M17.002 13.87L15.132 15l.426-2.204L14 11.146l2.11-.088L17.001 9l.892 2.058H20l-1.554 1.738l.468 2.204l-1.912-1.13Zm8 0L23.132 15l.426-2.204L22 11.146l2.11-.088L25.001 9l.892 2.058H28l-1.554 1.738l.468 2.204l-1.912-1.13Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackSt1"><use href="#flagpackSt0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}