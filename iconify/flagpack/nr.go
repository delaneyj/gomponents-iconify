package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackNr0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackNr2)"><use href="#flagpackNr0"/><path fill="#2E42A5" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackNr1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g fill-rule="evenodd" clip-rule="evenodd" mask="url(#flagpackNr1)"><path fill="#FECA00" d="M0 8v4h32V8H0Z"/><path fill="#F7FCFF" d="m8.83 19.58l-1.545 2.005l-.072-2.53l-2.428.714l1.43-2.09l-2.385-.85l2.384-.85l-1.43-2.088l2.43.714l.07-2.53L8.83 14.08l1.545-2.006l.071 2.53l2.429-.713l-1.43 2.089l2.385.85l-2.385.85l1.43 2.089l-2.429-.715l-.071 2.53L8.83 19.58Z"/></g></g><defs><clipPath id="flagpackNr2"><use href="#flagpackNr0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}