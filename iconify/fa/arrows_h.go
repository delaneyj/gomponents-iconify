package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1792 640q0 26-19 45l-256 256q-19 19-45 19t-45-19t-19-45V768H384v128q0 26-19 45t-45 19t-45-19L19 685Q0 666 0 640t19-45l256-256q19-19 45-19t45 19t19 45v128h1024V384q0-26 19-45t45-19t45 19l256 256q19 19 19 45z"/>`),
		g.Group(children),
	)
}