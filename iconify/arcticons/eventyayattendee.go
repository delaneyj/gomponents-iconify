package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eventyayattendee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.65 4.5h-9.31a4.34 4.34 0 0 1-8.68 0h-9.31a2 2 0 0 0-1.95 2v35.1a2 2 0 0 0 1.95 2h9.31a4.34 4.34 0 1 1 8.68 0h9.31a2 2 0 0 0 2-2V6.45a2 2 0 0 0-2-1.95ZM24 33.06A9.06 9.06 0 1 1 33.06 24A9.06 9.06 0 0 1 24 33.06Zm-3.92-7.18l7.84-3.76a4.35 4.35 0 1 0-2 5.8h0a4.35 4.35 0 0 0 2.4-3.16"/>`),
		g.Group(children),
	)
}