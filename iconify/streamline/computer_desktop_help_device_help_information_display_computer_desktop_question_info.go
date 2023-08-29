package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerDesktopHelpDeviceHelpInformationDisplayComputerDesktopQuestionInfo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 2h1a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-.5.5H1a.5.5 0 0 1-.5-.5v-8A.5.5 0 0 1 1 2h1m4 9l-1 2.5M8 11l1 2.5m-5 0h6"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5 2.5a2 2 0 1 1 2 2v1"/><path fill="currentColor" d="M7 7.5a.75.75 0 1 0 .75.75A.75.75 0 0 0 7 7.5Z"/>`),
		g.Group(children),
	)
}