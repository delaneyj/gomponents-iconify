package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Et(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackEt0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackEt2)"><use href="#flagpackEt0"/><path fill="#FECA00" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackEt1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g mask="url(#flagpackEt1)"><path fill="#5EAA22" fill-rule="evenodd" d="M0 0v8h32V0H0Z" clip-rule="evenodd"/><path fill="#E31D1C" fill-rule="evenodd" d="M0 16v8h32v-8H0Z" clip-rule="evenodd"/><path fill="#2B77B8" fill-rule="evenodd" d="M16 18a6 6 0 1 0 0-12a6 6 0 0 0 0 12Z" clip-rule="evenodd"/><path stroke="#FECA00" stroke-width="1.5" d="m16 14l-2.762.927l.86-2.309l-1.874-2.236h2.6L16 8l1.176 2.382h2.657l-1.93 2.236l.684 2.309L16 14Z" clip-rule="evenodd"/><path stroke="#2B77B8" d="m15.696 12.034l-2.085 4.36m1.757-4.914h-4m4.754 1.358l3.909 1.804m-3.064-2.768l2.83-3.098"/></g></g><defs><clipPath id="flagpackEt2"><use href="#flagpackEt0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}