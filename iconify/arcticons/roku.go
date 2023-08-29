package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Roku(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2h33a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10 28.3v-9.54h3.12a3.2 3.2 0 1 1 0 6.4H10m3.12 0l3.13 3.13"/><rect width="4.77" height="6.32" x="18.64" y="21.98" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.39"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.23 22v3.93a2.39 2.39 0 0 0 2.38 2.39h0A2.39 2.39 0 0 0 38 25.91V22m0 3.91v2.39m-11.86-9.54v9.54m0-2.03l4.32-4.3m-2.94 2.93l3.39 3.38"/>`),
		g.Group(children),
	)
}