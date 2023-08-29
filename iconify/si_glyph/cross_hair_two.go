package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrossHairTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.979 10c1.093 0 1.979-.962 1.979-2.043c0-1.08-.886-1.957-1.979-1.957A1.968 1.968 0 0 0 7 7.957C7 9.037 7.886 10 8.979 10zM8 0h1.986v4H8zm0 12v4.003h2.007V12H8zM1 7h3.984v1.969H1zm12 0v1.928h3.99V7H13z"/>`),
		g.Group(children),
	)
}