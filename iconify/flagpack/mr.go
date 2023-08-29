package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackMr0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackMr2)"><use href="#flagpackMr0"/><path fill="#1C7B4D" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackMr1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g fill-rule="evenodd" clip-rule="evenodd" mask="url(#flagpackMr1)"><path fill="#E31D1C" d="M0 0v6h32V0H0Zm0 18v6h32v-6H0Z"/><path fill="#FECA00" d="M16.242 14.379c5.185.028 6.646-4.314 6.646-4.314c-.294 3.715-2.349 6.234-6.646 6.234s-5.875-3.263-6.646-6.518c0 0 1.46 4.569 6.646 4.598Z"/><path fill="#FECA00" d="m17.637 9.814l.337 1.963l-1.764-.927l-1.763.927l.337-1.963l-1.427-1.534h1.972l.881-1.93l.882 1.93h1.972l-1.427 1.534Z"/></g></g><defs><clipPath id="flagpackMr2"><use href="#flagpackMr0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}