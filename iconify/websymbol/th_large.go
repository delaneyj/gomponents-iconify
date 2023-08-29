package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThLarge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M1000 0v440H560V0h440zM680 320h200V120H680v200zM440 0v440H0V0h440zM120 320h200V120H120v200zm880 240v440H560V560h440zM680 880h200V680H680v200zM440 560v440H0V560h440zM120 880h200V680H120v200z"/>`),
		g.Group(children),
	)
}