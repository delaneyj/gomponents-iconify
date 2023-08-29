package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deluge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 4.5l11 18.32a14.43 14.43 0 0 1 2.29 7.44a13.24 13.24 0 0 1-26.48 0a14.43 14.43 0 0 1 2.29-7.44L24 4.5ZM30.1 42a4.71 4.71 0 0 0 1.19-4.36c-.84-3.67-4.41-4.31-8-4.31c-3.26 0-5.11-2.42-5.11-5.53c0-6.57 6.82-10.31 6.82-10.31h0s-9.51 1.31-12.85 6.9"/>`),
		g.Group(children),
	)
}