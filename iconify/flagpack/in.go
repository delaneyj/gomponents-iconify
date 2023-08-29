package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func In(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackIn0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackIn2)"><use href="#flagpackIn0"/><path fill="#F7FCFF" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackIn1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g fill-rule="evenodd" clip-rule="evenodd" mask="url(#flagpackIn1)"><path fill="#FF8C1A" d="M0 0v8h32V0H0Z"/><path fill="#5EAA22" d="M0 16v8h32v-8H0Z"/><path fill="#3D58DB" d="M12 12a4 4 0 1 0 8 0a4 4 0 0 0-8 0Zm7 0a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z"/><path fill="#3D58DB" d="m15.995 12.86l-.571 3.121l.331-3.155l-1.427 2.834l1.207-2.934l-2.167 2.316l1.984-2.474l-2.732 1.612l2.602-1.816l-3.076.777l3.007-1.009l-3.17-.121l3.17-.12l-3.007-1.01l3.076.777l-2.602-1.815l2.732 1.612l-1.984-2.475l2.167 2.316l-1.207-2.934l1.427 2.834l-.331-3.155l.57 3.12l.571-3.12l-.331 3.155l1.427-2.834l-1.207 2.934L18.62 8.98l-1.984 2.475l2.732-1.612l-2.602 1.815l3.076-.777l-3.008 1.01l3.17.12l-3.17.121l3.008 1.01l-3.076-.778l2.602 1.816l-2.732-1.612l1.984 2.474l-2.167-2.316l1.207 2.934l-1.427-2.834l.331 3.155l-.57-3.12Z"/></g></g><defs><clipPath id="flagpackIn2"><use href="#flagpackIn0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}