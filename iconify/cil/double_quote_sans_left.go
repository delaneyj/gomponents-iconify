package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleQuoteSansLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M232 246.857V16H16v400h38.4ZM48 48h152v185.143L48 377.905Zm232 368h38.4L496 246.857V16H280Zm32-368h152v185.143L312 377.905Z"/>`),
		g.Group(children),
	)
}