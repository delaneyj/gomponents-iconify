package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EntertainmentControlButtonStopButtonTelevisionButtonsMoviesStopTvVideoControls(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<rect width="13" height="13" x=".5" y=".52" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/>`),
		g.Group(children),
	)
}