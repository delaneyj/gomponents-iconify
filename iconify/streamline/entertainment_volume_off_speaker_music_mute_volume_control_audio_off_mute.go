package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EntertainmentVolumeOffSpeakerMusicMuteVolumeControlAudioOffMute(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.5 13.5l13-13M4.5 5H3a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h1.5ZM10 4V3a1 1 0 0 0-.55-.89a1 1 0 0 0-1 .08L4.5 5m2.17 5.56l1.74 1.25a1 1 0 0 0 1 .08A1 1 0 0 0 10 11V7M4.5 9l.29.21"/>`),
		g.Group(children),
	)
}