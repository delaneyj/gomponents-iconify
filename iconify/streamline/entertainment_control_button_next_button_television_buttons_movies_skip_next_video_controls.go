package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EntertainmentControlButtonNextButtonTelevisionButtonsMoviesSkipNextVideoControls(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5.5v13m-13-1.84a1 1 0 0 0 .52.88a1 1 0 0 0 1 0l7.19-4.7a1 1 0 0 0 0-1.68L2 1.5a1 1 0 0 0-1 0a1 1 0 0 0-.52.88Z"/>`),
		g.Group(children),
	)
}