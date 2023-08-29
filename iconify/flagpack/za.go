package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Za(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackZa0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackZa2)"><use href="#flagpackZa0"/><path fill="#F7FCFF" fill-rule="evenodd" d="M0 0h32v24H0V0Z" clip-rule="evenodd"/><path fill="#E31D1C" fill-rule="evenodd" d="M0 0v8h32V0H0Z" clip-rule="evenodd"/><path fill="#3D58DB" fill-rule="evenodd" d="M0 16v8h32v-8H0Z" clip-rule="evenodd"/><mask id="flagpackZa1" width="36" height="38" x="-2" y="-7" fill="#000" maskUnits="userSpaceOnUse"><path fill="#fff" d="M-2-7h36v38H-2z"/><path fill-rule="evenodd" d="M15.429 10L0-2v28l15.429-12H32v-4H15.429Z" clip-rule="evenodd"/></mask><path fill="#5EAA22" fill-rule="evenodd" d="M15.429 10L0-2v28l15.429-12H32v-4H15.429Z" clip-rule="evenodd"/><path fill="#F7FCFF" d="m0-2l1.228-1.579L-2-6.089V-2h2Zm15.429 12L14.2 11.579l.541.421h.687v-2ZM0 26h-2v4.09l3.228-2.511L0 26Zm15.429-12v-2h-.687l-.541.421L15.429 14ZM32 14v2h2v-2h-2Zm0-4h2V8h-2v2ZM-1.228-.421l15.429 12l2.456-3.158l-15.43-12L-1.227-.42ZM2 26V-2h-4v28h4Zm12.2-13.579l-15.428 12l2.456 3.158l15.429-12L14.2 12.42ZM32 12H15.429v4H32v-4Zm-2-2v4h4v-4h-4Zm-14.571 2H32V8H15.429v4Z" mask="url(#flagpackZa1)"/><path fill="#272727" stroke="#FECA00" stroke-width="2" d="M.6 5.2L-1 4v16l1.6-1.2l8-6l1.067-.8l-1.067-.8l-8-6Z"/></g><defs><clipPath id="flagpackZa2"><use href="#flagpackZa0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}