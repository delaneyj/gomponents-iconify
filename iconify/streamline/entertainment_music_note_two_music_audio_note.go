package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EntertainmentMusicNoteTwoMusicAudioNote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="2.5" cy="11.42" r="2"/><circle cx="11.5" cy="8.92" r="2"/><path d="M13.5 8.92V1.08a.5.5 0 0 0-.63-.48l-8 2.22a.5.5 0 0 0-.37.48v8.12m0-5.5l9-2.5"/></g>`),
		g.Group(children),
	)
}