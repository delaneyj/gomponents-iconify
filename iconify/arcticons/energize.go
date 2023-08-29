package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Energize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.365 16.033L34.755 9.5c1.197 2.817-.358 5.497-3.207 6.707l-15.39 6.532c-1.196-2.816.382-5.507 3.207-6.706ZM14.662 26.23l15.39-6.532c1.197 2.817-.358 5.497-3.207 6.707l-15.39 6.532c-1.196-2.816.381-5.507 3.207-6.706Zm6.49 5.563l15.39-6.532c1.196 2.817-.358 5.497-3.207 6.706L17.945 38.5c-1.196-2.816.381-5.507 3.207-6.707Z"/>`),
		g.Group(children),
	)
}