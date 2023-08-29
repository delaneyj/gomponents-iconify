package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tapacademy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.463 22.987v-6.032h1.975c1.117 0 2.022.907 2.022 2.026s-.905 2.025-2.022 2.025h-1.975M5.5 16.955h3.996m-1.998 6.032v-6.032m2.905 6.014l1.959-6.014m1.961 6.032l-1.961-6.032m1.305 4.014h-2.613m15.599 10.076h4.917m-4.917-6.032h3.016m-3.016 3.016h1.967m2.95 1.508v-4.524l3.016 6.032l3.016-6.023v6.023M24.9 27.652v.754a2.639 2.639 0 0 1-2.64 2.64h-1.357v-6.033H24.9m18.698 0L41.6 28.029l-1.998-3.016m1.998 6.032v-3.016m-27.241.993v.025a1.998 1.998 0 0 1-1.998 1.998h0a1.998 1.998 0 0 1-1.998-1.998v-2.036c0-1.103.894-1.998 1.998-1.998h0c1.103 0 1.998.895 1.998 1.998v.025M5.5 31.027l1.959-6.014m1.962 6.032l-1.962-6.032m1.306 4.014H6.152m9.169 2l1.959-6.014m1.962 6.032l-1.962-6.032m1.306 4.014h-2.614"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}