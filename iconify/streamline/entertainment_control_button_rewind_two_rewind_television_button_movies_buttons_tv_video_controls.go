package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EntertainmentControlButtonRewindTwoRewindTelevisionButtonMoviesButtonsTvVideoControls(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 10.29a.7.7 0 0 1-.37.62a.71.71 0 0 1-.73 0L7.32 7.59a.7.7 0 0 1 0-1.18l5.08-3.3a.71.71 0 0 1 .73 0a.7.7 0 0 1 .37.62Z"/><path d="M7 10.29a.7.7 0 0 1-.37.62a.71.71 0 0 1-.73 0L.82 7.59a.7.7 0 0 1 0-1.18l5.08-3.3a.71.71 0 0 1 .73 0a.7.7 0 0 1 .37.6ZM.5 2.5v9"/></g>`),
		g.Group(children),
	)
}