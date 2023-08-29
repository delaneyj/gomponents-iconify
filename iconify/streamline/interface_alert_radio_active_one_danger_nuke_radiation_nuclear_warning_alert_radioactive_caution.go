package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceAlertRadioActiveOneDangerNukeRadiationNuclearWarningAlertRadioactiveCaution(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6 4.83L3.75.94A6.49 6.49 0 0 0 .5 6.56H5a2 2 0 0 1 1-1.73Zm3 1.73h4.5A6.49 6.49 0 0 0 10.25.94L8 4.83a2 2 0 0 1 1 1.73Zm-2 2a2 2 0 0 1-1-.26l-2.25 3.89a6.51 6.51 0 0 0 6.5 0L8 8.3a2 2 0 0 1-1 .26Z"/>`),
		g.Group(children),
	)
}