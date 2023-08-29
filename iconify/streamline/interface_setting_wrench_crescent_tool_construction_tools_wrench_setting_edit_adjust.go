package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceSettingWrenchCrescentToolConstructionToolsWrenchSettingEditAdjust(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.43 3.59a.76.76 0 0 0-.35-.51l-2 2a1 1 0 0 1-1.44 0l-.76-.68a1 1 0 0 1 0-1.4l2-2a.76.76 0 0 0-.48-.43A3.8 3.8 0 0 0 6.26 6L.8 11.41a1 1 0 0 0 0 1.43l.36.36a1 1 0 0 0 1.43 0l5.46-5.45a3.81 3.81 0 0 0 5.38-4.16Z"/>`),
		g.Group(children),
	)
}