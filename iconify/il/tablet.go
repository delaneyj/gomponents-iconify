package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tablet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M510 8q19 0 33 13t13 34v648q0 20-13 33t-33 14H46q-19 0-32-14T0 703V55q0-20 14-34T46 8h464zM278 703q20 0 33-14t13-33t-13-32t-33-13t-33 13t-13 32.5t13 32.5t33 14zM486 89q0-11-11-11H81q-5 0-8 3t-4 8v464q0 5 4 8t8 3h394q11 0 11-11V89z"/>`),
		g.Group(children),
	)
}