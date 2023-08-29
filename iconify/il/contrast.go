package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Contrast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M371 8q76 0 144 29t118 79t79 118t29 145t-29 144t-79 118t-118 80t-144.5 29T227 721t-118-80t-80-118T0 379t29-145t80-118t118-79T371 8zm0 117q0-11-7-18t-19-5q-53 5-99 28t-80 61t-53 85t-20 103t20 102t53 86t80 60t99 29q11 1 19-6t7-17V125z"/>`),
		g.Group(children),
	)
}