package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Luca(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.462 17.871v9.625a1.3 1.3 0 0 0 1.375 1.375h.413m16.135-1.375a2.662 2.662 0 0 1-2.338 1.375h0a2.758 2.758 0 0 1-2.75-2.75v-1.787a2.758 2.758 0 0 1 2.75-2.75h0a2.662 2.662 0 0 1 2.338 1.375m-8.412 3.162a2.758 2.758 0 0 1-2.75 2.75h0a2.758 2.758 0 0 1-2.75-2.75v-4.537m6.6 7.287a1.183 1.183 0 0 1-1.1-1.1v-6.187m16.465 4.537a2.758 2.758 0 0 1-2.75 2.75h0a2.758 2.758 0 0 1-2.75-2.75v-1.787a2.758 2.758 0 0 1 2.75-2.75h0a2.758 2.758 0 0 1 2.75 2.75"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.538 28.871a1.183 1.183 0 0 1-1.1-1.1v-6.187"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2.006 2.006 0 0 1-2-2v-33a2.006 2.006 0 0 1 2-2h33a2.006 2.006 0 0 1 2 2v33a2.006 2.006 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}