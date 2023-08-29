package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Airplay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M464 47H48a32.036 32.036 0 0 0-32 32v297a32.036 32.036 0 0 0 32 32h76.448l24.89-32H48V79h416l.02 297H362.662l24.89 32H464a32.036 32.036 0 0 0 32-32V79a32.036 32.036 0 0 0-32-32Z"/><path fill="currentColor" d="M98.834 496h314.332L256 293.939Zm65.431-32L256 346.061L347.735 464Z"/>`),
		g.Group(children),
	)
}