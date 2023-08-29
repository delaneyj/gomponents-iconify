package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EdgeScreenNoteTen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M20.98 30.041v-8.08a4.002 4.002 0 0 1 4.002-4.002h0a4.002 4.002 0 0 1 4.002 4.002v8.08M32.459 24h3.939m2.102 6.041h-2.039a4.002 4.002 0 0 1-4.002-4.002V21.96a4.002 4.002 0 0 1 4.002-4.002H38.5"/><rect width="8.005" height="12.082" x="9.5" y="17.959" rx="4.002" ry="4.002"/></g>`),
		g.Group(children),
	)
}