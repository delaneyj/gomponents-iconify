package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EntertainmentControlButtonPreviousButtonTelevisionButtonsMoviesSkipPreviousVideoControls(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5.5v13m13-1.84a1 1 0 0 1-.52.88a1 1 0 0 1-1 0l-7.19-4.7a1 1 0 0 1 0-1.68L12 1.5a1 1 0 0 1 1 0a1 1 0 0 1 .52.88Z"/>`),
		g.Group(children),
	)
}