package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flixbus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.608 10.39v11.594h5.332m3.281-11.594v11.594M9.854 10.39h5.797m-5.797 5.797h3.768M9.854 10.39v11.594M38.146 10.39l-2.683 4.05m0 3.494l2.683 4.05M30.465 10.39l3.84 5.797l-3.84 5.797m.585 14.356a3.243 3.243 0 0 0 2.842 1.27h1.716a2.895 2.895 0 0 0 2.892-2.9v0a2.895 2.895 0 0 0-2.892-2.898h-1.896a2.895 2.895 0 0 1-2.892-2.898h0a2.895 2.895 0 0 1 2.892-2.899h1.716a3.243 3.243 0 0 1 2.843 1.27m-18.143-1.269v7.753a3.84 3.84 0 0 0 7.68 0v-7.753"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".866" d="M14.282 31.813a2.898 2.898 0 1 1 0 5.796H9.5V26.016h4.782a2.898 2.898 0 0 1 0 5.797Zm0 0H9.5"/>`),
		g.Group(children),
	)
}