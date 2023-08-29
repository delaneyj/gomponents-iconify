package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceRemoveTwoDeleteBoxSubtractButtonsButtonRemoveTextboxTextAddBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<rect width="13" height="3" x=".5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/>`),
		g.Group(children),
	)
}