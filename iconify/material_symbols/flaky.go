package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flaky(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22ZM7.35 11.225L8.75 9.8l1.4 1.425l1.075-1.075L9.8 8.75l1.425-1.4l-1.075-1.075L8.75 7.7l-1.4-1.425L6.275 7.35L7.7 8.75l-1.425 1.4l1.075 1.075ZM12 20q3.325 0 5.663-2.337T20 12q0-1.65-.625-3.1T17.65 6.35l-11.3 11.3q1.1 1.1 2.55 1.725T12 20Zm2.05-2.4l-2.4-2.4l1.05-1.075l1.35 1.35L16.55 13l1.05 1.05l-3.55 3.55Z"/>`),
		g.Group(children),
	)
}