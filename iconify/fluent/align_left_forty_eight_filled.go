package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeftFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.5 5.25a1.25 1.25 0 1 1 2.5 0v37.5a1.25 1.25 0 1 1-2.5 0V5.25ZM15.75 26a4.25 4.25 0 0 0-4.25 4.25v5.5A4.25 4.25 0 0 0 15.75 40h15A4.25 4.25 0 0 0 35 35.75v-5.5A4.25 4.25 0 0 0 30.75 26h-15ZM11.5 12.25v5.5A4.25 4.25 0 0 0 15.75 22H38a4.25 4.25 0 0 0 4.25-4.25v-5.5A4.25 4.25 0 0 0 38 8H15.75a4.25 4.25 0 0 0-4.25 4.25Z"/>`),
		g.Group(children),
	)
}