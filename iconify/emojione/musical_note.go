package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicalNote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#4d5357" d="M25.4 2v36.7c-1.2-.4-3.7-.6-5.2-.6c-13.6 0-13.6 16.6 0 16.6c5.9 0 11.7-3.7 11.7-8.3V25.8L47.4 31v15c-1.2-.4-3.7-.6-5.2-.6c-13.6 0-13.6 16.6 0 16.6c5.9 0 11.7-3.7 11.7-8.3V11.4L25.4 2m22 21.5l-15.5-5.3v-6.3l15.5 5.4v6.2"/>`),
		g.Group(children),
	)
}