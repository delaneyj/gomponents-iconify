package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Watch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g stroke-miterlimit="10" stroke-width="2"><path fill="#9b9b9a" stroke="#9b9b9a" d="M53.996 40.523h.699c1.1 0 2-.9 2-2V33.37c0-1.1-.9-2-2-2H50v9.165l3.996-.01z"/><path fill="#9b9b9a" d="M25 21V9a1 1 0 0 1 1-1h20a1 1 0 0 1 1 1v12"/><path fill="#9b9b9a" stroke="#9b9b9a" d="M47 51v12a1 1 0 0 1-1 1H26a1 1 0 0 1-1-1V51"/><path fill="#d0cfce" stroke="#d0cfce" d="M48 51H24c-1.1 0-2-.9-2-2V23c0-1.1.9-2 2-2h24c1.1 0 2 .9 2 2v26c0 1.1-.9 2-2 2z"/></g><g fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2"><path d="M25 21V9a1 1 0 0 1 1-1h20a1 1 0 0 1 1 1v12"/><path stroke-linecap="round" stroke-linejoin="round" d="M35.844 26.188v10h8.312M47 55v9H25v-9m0-38V8h22v9m6.996 14.37h.699c1.1 0 2 .9 2 2v5.153c0 1.1-.9 2-2 2h-.7"/><path stroke-linecap="round" stroke-linejoin="round" d="M48 51H24c-1.1 0-2-.9-2-2V23c0-1.1.9-2 2-2h24c1.1 0 2 .9 2 2v26c0 1.1-.9 2-2 2z"/><path stroke-linecap="round" stroke-linejoin="round" d="M35.844 26.188v10h8.312"/></g>`),
		g.Group(children),
	)
}