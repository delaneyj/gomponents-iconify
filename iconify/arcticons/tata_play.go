package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TataPlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 42.5h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.25 14.259h6.549m-3.274 7.715v-7.715m3.274 7.692l3.212-7.692m3.214 7.715l-3.214-7.715m3.264 0h6.549m-3.275 7.715v-7.715m3.275 7.692l3.211-7.692m3.215 7.715l-3.215-7.715M18.388 26.052v7.689h4.831M10.75 33.74v-7.688h3.19c1.422 0 2.576 1.156 2.576 2.583s-1.154 2.582-2.577 2.582H10.75m27-5.165l-3.2 3.845l-3.2-3.845m3.2 7.689v-3.844m-8.959 3.821l3.14-7.666l3.14 7.689"/>`),
		g.Group(children),
	)
}