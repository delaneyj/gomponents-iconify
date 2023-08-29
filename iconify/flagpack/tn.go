package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackTn0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackTn2)"><use href="#flagpackTn0"/><path fill="#E31D1C" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackTn1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g fill-rule="evenodd" clip-rule="evenodd" mask="url(#flagpackTn1)"><path fill="#F7FCFF" d="M16 20a8 8 0 1 0 0-16a8 8 0 0 0 0 16Z"/><path fill="#E31D1C" d="M17.403 17.65s-4.217-1.15-4.217-5.676c0-4.527 4.217-5.776 4.217-5.776c-1.744-.675-6.846.36-6.846 5.775c0 5.416 5.245 6.391 6.846 5.678Zm-.232-6.662l-2.092.765l2.248.786l.076 2.104l1.368-1.635l2.256.16l-1.625-1.326l.979-1.915l-1.913.644l-1.325-1.656l.028 2.073Z"/></g></g><defs><clipPath id="flagpackTn2"><use href="#flagpackTn0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}