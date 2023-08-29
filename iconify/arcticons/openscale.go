package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Openscale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.55 5.5H7.45a2 2 0 0 0-1.95 2v33.1a2 2 0 0 0 2 2h33.1a2 2 0 0 0 2-2V7.45a2 2 0 0 0-2.05-1.95Zm-4.1 12l-3.68 3.67l-1.22 1.22a2.09 2.09 0 0 1-3 0a6.47 6.47 0 0 0-9.16 0a2.09 2.09 0 0 1-3 0l-1.22-1.22l-3.68-3.67a2.12 2.12 0 0 1 0-3a17.6 17.6 0 0 1 24.9 0a2.12 2.12 0 0 1 .06 3.02ZM19 14.93l3.62 5.73"/>`),
		g.Group(children),
	)
}