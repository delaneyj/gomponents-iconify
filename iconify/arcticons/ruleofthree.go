package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ruleofthree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.56 20.91V11.5h3.08a3.16 3.16 0 0 1 0 6.32h-3.08m3.08 0l3.07 3.08m5.99-6.67l1.06-1.97l1.06 1.97m-4 5.5h5.88m-5.88-3.53h5.88m8.74 1a2.35 2.35 0 0 0-2.35-2.35h0a2.35 2.35 0 0 0-2.35 2.35v1.53a2.35 2.35 0 0 0 2.35 2.35h0a2.35 2.35 0 0 0 2.35-2.35m0 2.35v-9.4m0 18.59l-4.7 6.23m4.7 0l-4.7-6.23m-18.95-1.94v7A1.18 1.18 0 0 0 14 36.5h.35m-2.79-6.23h2.47m9.67-.62l1.06-1.97l1.06 1.97m-4 5.5h5.88m-5.88-3.53h5.88"/>`),
		g.Group(children),
	)
}