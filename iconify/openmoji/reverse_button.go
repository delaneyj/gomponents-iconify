package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReverseButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="none" stroke="#000" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M52.612 16.394a4.634 4.634 0 0 0-1.576-.297c-.848 0-1.697.297-2.424.772l-30 15.98l-.303.296c-.788.772-1.212 1.723-1.212 2.792c0 1.07.424 2.08 1.212 2.792l.303.297l30 16.098c1.09.832 2.667 1.01 4 .475c1.515-.594 2.485-2.079 2.485-3.683v-31.84c0-1.603-.97-3.088-2.485-3.682z"/>`),
		g.Group(children),
	)
}