package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Az(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackAz0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackAz2)"><use href="#flagpackAz0"/><path fill="#AF0100" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackAz1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g fill-rule="evenodd" clip-rule="evenodd" mask="url(#flagpackAz1)"><path fill="#3CA5D9" d="M0 0v8h32V0H0Z"/><path fill="#73BE4A" d="M0 16v8h32v-8H0Z"/><path fill="#F7FCFF" d="M17.14 15.024c-1.347-.31-2.53-1.47-2.516-3.024c.013-1.455.87-2.632 2.35-2.967c1.482-.334 3.018.301 3.018.301c-.408-.907-1.83-1.544-2.995-1.541c-2.17.006-4.486 1.663-4.509 4.193c-.023 2.623 2.473 4.114 4.67 4.108c1.761-.005 2.598-1.138 2.772-1.62c0 0-1.443.86-2.79.55Zm.878-1.603l1.175-.818l1.176.818l-.415-1.371l1.142-.865l-1.432-.03l-.47-1.352l-.47 1.352l-1.433.03l1.142.865l-.415 1.37Z"/></g></g><defs><clipPath id="flagpackAz2"><use href="#flagpackAz0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}