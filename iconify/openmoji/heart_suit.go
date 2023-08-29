package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartSuit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#EA5A47" d="M60.7 26.2c0-7.2-5.9-13.1-13.1-13.1c-5 0-9.3 2.8-11.5 6.9c-2.2-4.1-6.6-6.9-11.5-6.9c-7.2 0-13.1 5.9-13.1 13.1c0 3.1 1.1 6 2.9 8.2l21.8 27l21.8-27c1.6-2.2 2.7-5 2.7-8.2z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M60.7 26.2c0-7.2-5.9-13.1-13.1-13.1c-5 0-9.3 2.8-11.5 6.9c-2.2-4.1-6.6-6.9-11.5-6.9c-7.2 0-13.1 5.9-13.1 13.1c0 3.1 1.1 6 2.9 8.2h0l21.8 27l21.8-27h0c1.6-2.2 2.7-5 2.7-8.2z"/>`),
		g.Group(children),
	)
}