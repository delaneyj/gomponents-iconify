package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Magic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m20.875 2.563l-.688.75l-1.687 1.78h-3.594v3.5l-1.719 1.813l-.687.719l2.188 2.188L3.03 25l-.719.719l.72.687l3.28 3.282l.688-.72l11.688-11.655l2.187 2.187l.719-.688l1.812-1.718h3.5V13.5l1.782-1.688l.75-.687l-2.532-2.531v-3.5h-3.5zm.031 2.874l1.375 1.375l.313.282h2.312v2.312l.282.313l1.375 1.375l-1.344 1.281l-.313.281v2.438h-2.312l-.282.281l-1.406 1.344l-.812-.813l4.531-4.531l-3.969-3.969l-.718.688l-3.813 3.844l-.844-.844l1.344-1.406l.281-.282V7.094h2.438l.281-.313zm-.25 4.782l1.125 1.156l-15.468 15.5l-1.157-1.156zM19 21v1h-1v2h1v1h2v-1h1v-2h-1v-1zm6 2v2h-2v2h2v2h2v-2h2v-2h-2v-2z"/>`),
		g.Group(children),
	)
}