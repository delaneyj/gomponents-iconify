package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackHn0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackHn2)"><use href="#flagpackHn0"/><path fill="#F7FCFF" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackHn1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g fill="#4564F9" fill-rule="evenodd" clip-rule="evenodd" mask="url(#flagpackHn1)"><path d="M0 0v8h32V0H0Zm0 16v8h32v-8H0Zm16.402-3.33l-.866.448l.165-.947l-.701-.74h.969l.433-.931l.433.93h.97l-.702.74l.166.948l-.867-.447Zm-6.084-2.167l-.866.447l.165-.947l-.701-.74h.969l.433-.931l.433.93h.97l-.702.74l.166.948l-.867-.447Zm0 4.2l-.866.447l.165-.947l-.701-.74h.969l.433-.931l.433.93h.97l-.702.74l.166.948l-.867-.447Zm12-4.2l-.866.447l.165-.947l-.701-.74h.969l.433-.931l.433.93h.97l-.702.74l.166.948l-.867-.447Zm0 4.2l-.866.447l.165-.947l-.701-.74h.969l.433-.931l.433.93h.97l-.702.74l.166.948l-.867-.447Z"/></g></g><defs><clipPath id="flagpackHn2"><use href="#flagpackHn0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}