package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerWebcamVideoWorkVideoMeetingCameraCompanyConferenceOffice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.82 3.75L10 5V3.5a1 1 0 0 0-1-1H1.5a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1H9a1 1 0 0 0 1-1V9l2.82 1.25a.5.5 0 0 0 .68-.47V4.22a.5.5 0 0 0-.68-.47Z"/>`),
		g.Group(children),
	)
}