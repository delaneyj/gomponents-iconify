package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shuttle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m910 386l-12 12q-27 51-46.5 122.5T832 673v224q0 26-50.5 77t-77.5 51L473 823l-58 58q-11 11-23.5 14.5t-27.5-2t-27.5-12t-33-24.5t-32.5-30l-36.5-36.5L198 754l-30-32.5l-24.5-33l-12-27.5l-2-27.5L144 610l58-58L0 321q0-27 51-77.5t77-50.5h224q81 0 152.5-19.5T627 127l12-12q26-26 97-60T883 5t117 20q36 41 20 117t-50 147t-60 97zM768 129q-15 0-36 29.5T704 193q48 24 76 52t52 76q5-7 34.5-28t29.5-36q0-53-37.5-90.5T768 129zM160 801q44-1 54.5 9.5T224 865q-1 47-62.5 103.5T32 1025H0v-32q0-68 56.5-129.5T160 801z"/>`),
		g.Group(children),
	)
}