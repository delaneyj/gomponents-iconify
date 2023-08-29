package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wallpapers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.65 21.33l4.93 7.74L35 27a.86.86 0 0 1 .1-.16a1.79 1.79 0 0 1 .51-.48h.05l.1-.06h1.19l.18.05h.16l.15.07h0a2 2 0 0 1 .55.46l.09.13v.06s0 0 0 0l5.1 7.54a1.84 1.84 0 0 1-1.61 2.73H6.32a1.84 1.84 0 0 1-1.84-1.84a1.81 1.81 0 0 1 .23-.88l14.58-22.95c0-.06.07-.11.1-.16v0a1.8 1.8 0 0 1 .54-.49l.15-.08s0 0 0 0l.14-.06h1.27l.14.07a1.92 1.92 0 0 1 .66.61a.35.35 0 0 0 0 .08h0L24 14.32m.22.06l4.43 6.95"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.64 21.18a5.36 5.36 0 1 0-3.59-5.49"/>`),
		g.Group(children),
	)
}