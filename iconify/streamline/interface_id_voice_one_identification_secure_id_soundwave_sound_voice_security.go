package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceIdVoiceOneIdentificationSecureIdSoundwaveSoundVoiceSecurity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 1.5v11m5.2-9v7M3.1 5v4m10.4-7.5v11m-2.6-9v7M8.3 5v4"/>`),
		g.Group(children),
	)
}