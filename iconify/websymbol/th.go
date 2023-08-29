package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Th(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M1000 0v160H560V0h440zm0 280v160H560V280h440zm0 280v160H560V560h440zM440 0v440H0V0h440zM120 320h200V120H120v200zm880 520v160H560V840h440zM440 560v440H0V560h440zM120 880h200V680H120v200z"/>`),
		g.Group(children),
	)
}