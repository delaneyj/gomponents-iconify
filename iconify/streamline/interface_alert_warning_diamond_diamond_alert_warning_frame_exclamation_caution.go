package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceAlertWarningDiamondDiamondAlertWarningFrameExclamationCaution(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 3.5v3"/><circle cx="7" cy="9.5" r=".5"/><rect width="9.82" height="9.82" x="2.09" y="2.09" rx="1.07" transform="rotate(-45 7 7)"/></g>`),
		g.Group(children),
	)
}