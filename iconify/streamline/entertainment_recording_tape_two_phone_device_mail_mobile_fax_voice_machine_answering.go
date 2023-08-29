package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EntertainmentRecordingTapeTwoPhoneDeviceMailMobileFaxVoiceMachineAnswering(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="3" cy="7" r="2.5"/><circle cx="11" cy="7" r="2.5"/><path d="M3 9.5h8"/></g>`),
		g.Group(children),
	)
}