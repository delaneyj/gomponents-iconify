package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrightnessThreeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.675 19.1q2-1.05 3.163-2.95T17 12q0-2.25-1.163-4.15T12.675 4.9Q13.8 6.45 14.4 8.263T15 12q0 1.925-.6 3.738T12.675 19.1ZM9 22q-.775 0-1.525-.113T6 21.55q3.125-1.025 5.063-3.65T13 12q0-3.275-1.938-5.9T6 2.45q.725-.225 1.475-.338T9 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T19 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T9 22Zm6-10Z"/>`),
		g.Group(children),
	)
}