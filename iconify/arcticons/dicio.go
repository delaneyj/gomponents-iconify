package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dicio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.745 38.148v0s.176 4.63 9.315 4.339c9.138-.293 20.683-2.377 25.681-12.874C45.74 19.116 38.705 5.604 14.51 5.5c0 0-5.356 29.949 1.162 29.924c-6.324-.93-9.74.269-9.926 2.724Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.137 12.07h0c2.98.093 4.188 2.322 4.188 5.188v5.423c0 2.865-1.323 5.188-4.188 5.188h0c-2.866 0-4.19-2.323-4.19-5.188v-5.423c0-2.902 1.157-5.038 4.19-5.188Zm-.022 18.733v4.138"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.112 24.22a8.18 8.18 0 0 0 16.049 0"/>`),
		g.Group(children),
	)
}