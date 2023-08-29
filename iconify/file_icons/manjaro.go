package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Manjaro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 0v512h144V144h184V0H0zm184 184v328h144V184H184zM368 0v512h144V0H368z"/>`),
		g.Group(children),
	)
}