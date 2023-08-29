package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicNotesPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 56a8 8 0 0 1-8 8h-16v16a8 8 0 0 1-16 0V64h-16a8 8 0 0 1 0-16h16V32a8 8 0 0 1 16 0v16h16a8 8 0 0 1 8 8ZM88 118.25V204a36 36 0 1 1-16-29.92V64a8 8 0 0 1 6.06-7.76l56-14a8 8 0 0 1 3.88 15.52L88 70.25v31.5l70.06-17.51a8 8 0 0 1 3.88 15.52ZM72 204a20 20 0 1 0-20 20a20 20 0 0 0 20-20Zm144-84v52a36 36 0 1 1-16-29.92V120a8 8 0 0 1 16 0Zm-16 52a20 20 0 1 0-20 20a20 20 0 0 0 20-20Z"/>`),
		g.Group(children),
	)
}