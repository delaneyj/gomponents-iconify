package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceIdVoiceTwoIdentificationSecureIdSoundwaveSoundVoiceBracketsSecurity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 4V1.5a1 1 0 0 1 1-1H4M13.5 4V1.5a1 1 0 0 0-1-1H10M.5 10v2.5a1 1 0 0 0 1 1H4m9.5-3.5v2.5a1 1 0 0 1-1 1H10m-9-6h2.5l1.5-4L7 9l2-4l1.5 2.5H13"/>`),
		g.Group(children),
	)
}