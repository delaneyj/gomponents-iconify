package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EntertainmentControlButtonFastForwardOneButtonControlsFastForwardMoviesTelevisionVideoTv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5 10.29a.7.7 0 0 0 .37.62a.71.71 0 0 0 .73 0l5.08-3.3a.7.7 0 0 0 0-1.18L1.6 3.11a.71.71 0 0 0-.73 0a.7.7 0 0 0-.37.62Z"/><path d="M7 10.29a.7.7 0 0 0 .37.62a.71.71 0 0 0 .73 0l5.08-3.3a.7.7 0 0 0 0-1.18L8.1 3.11a.71.71 0 0 0-.73 0a.7.7 0 0 0-.37.6Z"/></g>`),
		g.Group(children),
	)
}