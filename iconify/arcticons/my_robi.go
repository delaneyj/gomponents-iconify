package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MyRobi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.49 35.7L24 27.9l-13.51-7.8h0L24 12.3l13.51 7.8v15.6L24 43.5l-13.51-7.8z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.51 35.7L24 27.9l13.51-7.8M24 12.3L10.49 4.5v15.6M24 12.3v31.2"/>`),
		g.Group(children),
	)
}