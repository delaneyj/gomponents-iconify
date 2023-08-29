package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.34 16.048L9.586 29.664a1.683 1.683 0 0 0 1.657 2.288h.687M9.896 19.824h4.806m8.368 8.498a5.128 5.128 0 0 1-4.63 3.63h0a2.67 2.67 0 0 1-2.629-3.63l.65-2.36a5.128 5.128 0 0 1 4.631-3.63h0a2.67 2.67 0 0 1 2.63 3.63"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.522 31.952a1.068 1.068 0 0 1-1.051-1.452l2.251-8.167m3.884-2.509a1.347 1.347 0 0 1 1.326 1.83l-2.84 10.298"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.077 31.952a1.347 1.347 0 0 1-1.326-1.83l1.577-5.721c.697-2.528-.787-4.577-3.315-4.577h0a6.466 6.466 0 0 0-5.838 4.577m-14.473-4.577h13.904"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2.006 2.006 0 0 1-2-2v-33a2.006 2.006 0 0 1 2-2h33a2.006 2.006 0 0 1 2 2v33a2.006 2.006 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}