package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackBi0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackBi2)"><use href="#flagpackBi0"/><rect width="32" height="24" fill="#5EAA22" rx="2"/><mask id="flagpackBi1" width="48" height="30" x="-8" y="-3" fill="#000" maskUnits="userSpaceOnUse"><path fill="#fff" d="M-8-3h48v30H-8z"/><path fill-rule="evenodd" d="M16 12L32 0H0l16 12Zm0 0L0 24h32L16 12Z" clip-rule="evenodd"/></mask><path fill="#DD2C2B" fill-rule="evenodd" d="M16 12L32 0H0l16 12Zm0 0L0 24h32L16 12Z" clip-rule="evenodd"/><path fill="#fff" d="m32 0l1.5 2l6-4.5H32V0ZM0 0v-2.5h-7.5l6 4.5L0 0Zm0 24l-1.5-2l-6 4.5H0V24Zm32 0v2.5h7.5l-6-4.5l-1.5 2ZM30.5-2l-16 12l3 4l16-12l-3-4ZM0 2.5h32v-5H0v5ZM17.5 10L1.5-2l-3 4l16 12l3-4Zm-3 0l-16 12l3 4l16-12l-3-4ZM0 26.5h32v-5H0v5ZM33.5 22l-16-12l-3 4l16 12l3-4Z" mask="url(#flagpackBi1)"/><path fill="#fff" fill-rule="evenodd" d="M16 18a6 6 0 1 0 0-12a6 6 0 0 0 0 12Z" clip-rule="evenodd"/><path fill="#DD2C2B" fill-rule="evenodd" stroke="#5EAA22" stroke-width=".25" d="m15.43 10.387l-1.162.013l.592-1l-.592-1l1.162.013L16 7.4l.57 1.013l1.162-.013l-.592 1l.592 1l-1.162-.013L16 11.4l-.57-1.013Zm-2.5 4.1l-1.162.013l.592-1l-.592-1l1.162.013l.57-1.013l.57 1.013l1.162-.013l-.592 1l.592 1l-1.162-.013l-.57 1.013l-.57-1.013Zm5 0l-1.162.013l.592-1l-.592-1l1.162.013l.57-1.013l.57 1.013l1.162-.013l-.592 1l.592 1l-1.162-.013l-.57 1.013l-.57-1.013Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackBi2"><use href="#flagpackBi0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}