package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Metro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.59 40s4.51-5.36 4.86-11.67c.46-8-2.52-12.86-6.48-16.3h0a22 22 0 0 0-6.49-4L24 40L17.51 8A20.74 20.74 0 0 0 11 12c-3.94 3.46-6.94 8.36-6.49 16.3c.4 6.36 4.9 11.7 4.9 11.7"/>`),
		g.Group(children),
	)
}