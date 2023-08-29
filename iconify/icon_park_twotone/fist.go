package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFist0"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M34 41c-5.83 3.175-12.5 2.5-16 1S6 35.5 6 27c0-4.247 5.771-6.33 7.51-6.86a.686.686 0 0 0 .49-.652V12a3.5 3.5 0 1 1 7 0V9.5a3.5 3.5 0 0 1 7 0V12a3.5 3.5 0 1 1 7 0v4.493a3.5 3.5 0 0 1 7 .003V30c0 4-3.87 8.751-8 11Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFist0)"/>`),
		g.Group(children),
	)
}