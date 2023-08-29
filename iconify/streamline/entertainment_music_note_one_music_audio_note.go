package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EntertainmentMusicNoteOneMusicAudioNote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="4.25" cy="11" r="2.5"/><path d="M6.75 11V.5h0a5.5 5.5 0 0 1 5.5 5.5h0"/></g>`),
		g.Group(children),
	)
}