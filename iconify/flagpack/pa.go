package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackPa0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackPa2)"><use href="#flagpackPa0"/><path fill="#F7FCFF" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackPa1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g fill-rule="evenodd" clip-rule="evenodd" mask="url(#flagpackPa1)"><path fill="#E31D1C" d="M16 0v12h16V0H16Zm7.02 19.15l-2.302 1.424l.894-2.391l-1.957-1.811h2.374l.992-2.587l.757 2.587h2.377l-1.713 1.81l.839 2.392l-2.26-1.424Z"/><path fill="#2E42A5" d="M9.02 8.365L6.719 9.79l.894-2.392l-1.957-1.81H8.03L9.021 3l.757 2.587h2.377l-1.713 1.81l.839 2.393l-2.26-1.425ZM0 12v12h16V12H0Z"/></g></g><defs><clipPath id="flagpackPa2"><use href="#flagpackPa0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}