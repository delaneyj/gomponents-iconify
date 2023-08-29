package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><rect width="6.03" height="7.98" x="21.51" y="21.79" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.01"/><circle cx="34.18" cy="29.52" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30 17.73v10.54a1.51 1.51 0 0 0 1.51 1.51H32m-18.93-4.97a3 3 0 0 1 3-3h0a3 3 0 0 1 3 3v2a3 3 0 0 1-3 3h0a3 3 0 0 1-3-3m0 2.97V17.73"/>`),
		g.Group(children),
	)
}