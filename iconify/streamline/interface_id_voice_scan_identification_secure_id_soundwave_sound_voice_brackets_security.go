package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceIdVoiceScanIdentificationSecureIdSoundwaveSoundVoiceBracketsSecurity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 5.18v4.14m4-4.14v4.14M5 6.07v2.36m6-3.25v4.14M9 6.07v2.36M.5 4V1.5a1 1 0 0 1 1-1H4M13.5 4V1.5a1 1 0 0 0-1-1H10M.5 10v2.5a1 1 0 0 0 1 1H4m9.5-3.5v2.5a1 1 0 0 1-1 1H10"/>`),
		g.Group(children),
	)
}