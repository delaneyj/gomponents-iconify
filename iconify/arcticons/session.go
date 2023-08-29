package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Session(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.51 4.5h-19a9 9 0 0 0-9 9h0a9.05 9.05 0 0 0 4.3 7.7L24 29.41V18.59h10.28a7.19 7.19 0 0 0 7.1-5.43a7.07 7.07 0 0 0-6.87-8.66Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.49 43.5h19a9 9 0 0 0 9-9h0a9.05 9.05 0 0 0-4.3-7.7L24 18.59v10.82H13.72a7.19 7.19 0 0 0-7.1 5.43a7.07 7.07 0 0 0 6.87 8.66Z"/>`),
		g.Group(children),
	)
}