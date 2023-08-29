package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Opensudoku(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linejoin="round" rx="1.95"/><path fill="none" stroke="currentColor" stroke-linejoin="round" d="M13.92 5.5v37m20.16-37v37m8.42-28.58h-37m37 20.16h-37"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.23 30.5a4.53 4.53 0 0 0 4.13-4.6v-4A4.36 4.36 0 0 0 24 17.5h0a4.36 4.36 0 0 0-4.36 4.36"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.64 21.86A4.36 4.36 0 0 0 24 26.22h0a4.36 4.36 0 0 0 4.36-4.36m-7.95 7.22a3.63 3.63 0 0 0 3.18 1.42h.64"/>`),
		g.Group(children),
	)
}