package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M62.7 85.3H20V512h384v-42.7H62.7v-384zM361.3 0v128h128L361.3 0zM340 0H105.3v426.7h384V149.3H340V0zm64 298.7h-64V384h-85.3v-85.3h-64L297.3 192L404 298.7z"/>`),
		g.Group(children),
	)
}