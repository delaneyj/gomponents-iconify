package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zooforzotero(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2h33a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2Z"/><rect width="5.69" height="7.54" x="21.78" y="15.34" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.84"/><rect width="5.69" height="7.54" x="30.26" y="15.34" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.84"/><rect width="5.69" height="7.54" x="30.26" y="28.96" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.84"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.59 31.81A2.86 2.86 0 0 1 28.44 29h0m-2.85-.04v7.54m-13.53-25h7.53l-7.53 11.37h7.53m3.06 12.19a2.84 2.84 0 0 1-2.47 1.44h0a2.84 2.84 0 0 1-2.84-2.84v-1.85A2.85 2.85 0 0 1 20.18 29h0A2.86 2.86 0 0 1 23 31.81v.92h-5.66m-3.79-6.11v8.46A1.43 1.43 0 0 0 15 36.5h.43m-3.37-7.54h2.98"/>`),
		g.Group(children),
	)
}