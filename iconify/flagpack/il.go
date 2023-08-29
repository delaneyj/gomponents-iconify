package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Il(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackIl0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackIl2)"><use href="#flagpackIl0"/><path fill="#F7FCFF" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackIl1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g fill="#3D58DB" fill-rule="evenodd" clip-rule="evenodd" mask="url(#flagpackIl1)"><path d="M0 3v3h32V3H0Zm0 15v3h32v-3H0Zm11.381-3.061h9.377L16.116 6.62l-4.736 8.32Zm7.772-1.01h-6.132l3.086-5.47l3.046 5.47Z"/><path d="M11.264 9.076h9.417l-4.545 8.085l-4.872-8.085Zm7.781.974h-5.994l3.058 5.481l2.936-5.481Z"/></g></g><defs><clipPath id="flagpackIl2"><use href="#flagpackIl0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}