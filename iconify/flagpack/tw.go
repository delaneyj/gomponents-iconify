package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackTw0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackTw2)"><use href="#flagpackTw0"/><path fill="#EF0000" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackTw1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g fill-rule="evenodd" clip-rule="evenodd" mask="url(#flagpackTw1)"><path fill="#2E42A5" d="M0 0v14h18V0H0Z"/><path fill="#FEFFFF" d="m8.73 10.81l-1.482 1.85l-.359-2.342l-2.207.86l.86-2.208L3.2 8.611L5.048 7.13L3.2 5.648l2.342-.359l-.86-2.207l2.207.86l.359-2.342L8.73 3.448L10.21 1.6l.36 2.342l2.207-.86l-.86 2.207l2.341.359L12.41 7.13l1.848 1.481l-2.341.36l.86 2.207l-2.208-.86l-.359 2.341l-1.481-1.848Zm0-.818a2.862 2.862 0 1 0 0-5.725a2.862 2.862 0 0 0 0 5.725Zm2.29-2.862a2.29 2.29 0 1 1-4.58 0a2.29 2.29 0 0 1 4.58 0Z"/></g></g><defs><clipPath id="flagpackTw2"><use href="#flagpackTw0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}