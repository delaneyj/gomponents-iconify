package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoteControlAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M59 135q44-44 106.5-44T272 135l-31 30q-31-31-75.5-31T89 165zM165.5 5Q262 5 331 74l-30 30q-56-56-135.5-56T30 104L0 74Q69 5 165.5 5zM226 198q10 0 17.5 7t6.5 17v207q0 10-7 17t-17 7H104q-10 0-17-7t-7-17V222q0-10 7-17.5t17-7.5zm3 213V240H101v171h128z"/>`),
		g.Group(children),
	)
}