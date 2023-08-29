package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EntertainmentRecordingTapeOneFilmTelevisionTvMoviesReelVideo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="5.5" cy="5.5" r="5"/><circle cx="5.5" cy="5.5" r="1.5"/><path d="M10.5 5.5V12a1.5 1.5 0 0 0 1.5 1.5h0a1.5 1.5 0 0 0 1.5-1.5v-.5"/></g>`),
		g.Group(children),
	)
}