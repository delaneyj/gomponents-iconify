package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrightnessThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 22q-.8 0-1.538-.113T6 21.55q3.15-1.025 5.075-3.65T13 12q0-3.275-1.925-5.9T6 2.45q.725-.225 1.463-.338T9 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T19 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T9 22Z"/>`),
		g.Group(children),
	)
}