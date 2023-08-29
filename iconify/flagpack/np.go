package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Np(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackNp0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackNp2)"><use href="#flagpackNp0"/><use href="#flagpackNp0"/><path fill="#C51918" stroke="#4857A1" d="m11.616 12.32l9.317 11.18H.5V.842L20.04 11.5h-9.109l.684.82Z"/><mask id="flagpackNp1" width="22" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" stroke="#fff" d="m11.616 12.32l9.317 11.18H.5V.842L20.04 11.5h-9.109l.684.82Z"/></mask><g fill-rule="evenodd" clip-rule="evenodd" mask="url(#flagpackNp1)"><path fill="#F7FCFF" d="M5.83 20.01L4.604 21.6l-.056-2.006l-1.926.566l1.134-1.657l-1.891-.673l1.89-.674l-1.133-1.657l1.926.566l.056-2.006l1.226 1.59l1.225-1.59l.056 2.007l1.926-.567l-1.134 1.657l1.891.674l-1.89.674l1.133 1.656l-1.926-.566l-.056 2.006l-1.225-1.59ZM5.8 8.139l-.64.83l-.03-1.048l-1.005.296l.592-.865L3.73 7l.987-.352l-.592-.865l1.006.296l.03-1.048l.64.83l.639-.83l.03 1.048l1.005-.296l-.592.865l.988.352l-.988.352l.592.865l-1.006-.296l-.03 1.048l-.639-.83Z"/><path fill="#F9FAFA" d="M5.666 7.894C7.892 7.903 9.08 6.64 9.08 6.64c.235 1.441-1.578 2.4-3.391 2.4c-1.814 0-2.912-1.307-3.588-2.4c0 0 1.34 1.245 3.566 1.254Z"/></g></g><defs><clipPath id="flagpackNp2"><use href="#flagpackNp0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}