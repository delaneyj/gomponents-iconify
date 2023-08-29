package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EntertainmentCameraVideoFilmTelevisionTvCameraMoviesVideoRecorder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="3.5" cy="3.5" r="3"/><circle cx="10.5" cy="4.5" r="2"/><rect width="7.5" height="4.5" x="3.5" y="9" rx="1"/><path d="M13.5 10v2.5"/></g>`),
		g.Group(children),
	)
}