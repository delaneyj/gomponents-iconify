package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dj(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackDj0" fill="#fff" d="M0 0h32v24H0z"/><path id="flagpackDj1" fill="#fff" fill-rule="evenodd" d="M0 0v24l18-12L0 0Z" clip-rule="evenodd"/></defs><g fill="none"><g clip-path="url(#flagpackDj4)"><use href="#flagpackDj0"/><path fill="#73BE4A" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackDj2" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g mask="url(#flagpackDj2)"><path fill="#69F" fill-rule="evenodd" d="M0-2v14h32V-2H0Z" clip-rule="evenodd"/></g><use href="#flagpackDj1" fill-rule="evenodd" clip-rule="evenodd"/><mask id="flagpackDj3" width="18" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><use href="#flagpackDj1" fill-rule="evenodd" clip-rule="evenodd"/></mask><g mask="url(#flagpackDj3)"><path fill="#E31D1C" fill-rule="evenodd" d="m7.002 14.07l-1.87 1.13l.426-2.204L4 11.347l2.11-.09l.892-2.057l.892 2.058H10l-1.554 1.738l.468 2.204l-1.912-1.13Z" clip-rule="evenodd"/></g></g><defs><clipPath id="flagpackDj4"><use href="#flagpackDj0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}