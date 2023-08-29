package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EntertainmentVolumeLevelHighSpeakerHighVolumeControlAudioMusic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 5H1.5a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1H3Zm0 4l3.91 2.81a1 1 0 0 0 1 .08A1 1 0 0 0 8.5 11V3a1 1 0 0 0-.5-.89a1 1 0 0 0-1 .08L3 5m9.5-1a4.38 4.38 0 0 1 1 3a6.92 6.92 0 0 1-1 3.5m-2-5A2.19 2.19 0 0 1 11 7a2.19 2.19 0 0 1-.5 1.5"/>`),
		g.Group(children),
	)
}