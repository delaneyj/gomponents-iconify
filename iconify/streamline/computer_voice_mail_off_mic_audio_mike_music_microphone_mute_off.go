package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerVoiceMailOffMicAudioMikeMusicMicrophoneMuteOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.5 13.5l13-13m-8 8a2.49 2.49 0 0 1-1-2V3a2.5 2.5 0 0 1 5 0v1.5m-.23 3.04A2.5 2.5 0 0 1 8 8.79m-4.42 1.63A4.46 4.46 0 0 1 2 7h0m10 0h0a4.49 4.49 0 0 1-4.5 4.5h-1a5.25 5.25 0 0 1-.56 0m1.06 0v2"/>`),
		g.Group(children),
	)
}