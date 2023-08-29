package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Iherb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM13.046 19.572v9.384m6.217-9.384v9.384m-6.217-4.71h6.217"/><circle cx="10.321" cy="19.865" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.321 22.739v6.217m18.67-3.889a2.346 2.346 0 0 1 2.346-2.346h0m-2.346 0v6.217m-2.751-1.184a2.345 2.345 0 0 1-2.04 1.184h0a2.346 2.346 0 0 1-2.345-2.346v-1.525a2.346 2.346 0 0 1 2.346-2.346h0a2.346 2.346 0 0 1 2.346 2.346v.763h-4.692m11.953-.763a2.346 2.346 0 0 1 2.346-2.346h0a2.346 2.346 0 0 1 2.346 2.346v1.525a2.346 2.346 0 0 1-2.346 2.346h0a2.346 2.346 0 0 1-2.346-2.346m0 2.346v-9.383"/>`),
		g.Group(children),
	)
}